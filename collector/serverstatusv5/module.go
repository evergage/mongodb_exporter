package serverstatusv5

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type metricDefinition struct {
	SourcePath    string
	PathSegments  []string
	Family        string
	Help          string
	Type          string
	Conversion    float64
	FixedLabels   map[string]string
	ValueLabel    string
	AllowedValues []string
	Transform     string
}
type dynamicFamily struct {
	Name    string
	Root    []string
	Allowed map[string]struct{}
}
type FamilyError struct {
	Family string
	Path   string
	Err    error
}
type Report struct {
	Errors  []FamilyError
	Emitted int
}
type familyDescriptor struct {
	desc       *prometheus.Desc
	labelNames []string
}
type sample struct {
	definition metricDefinition
	value      float64
	labels     map[string]string
}
type snapshotValue struct {
	seconds   float64
	increment float64
}

type Module struct {
	descriptors  map[string]familyDescriptor
	decodeErrors *prometheus.CounterVec
	unknownKeys  *prometheus.CounterVec
}

func New() *Module {
	schemas := make(map[string]map[string]bool)
	helps := make(map[string]string)
	for _, d := range metricDefinitions {
		labels := schemas[d.Family]
		if labels == nil {
			labels = make(map[string]bool)
			schemas[d.Family] = labels
			helps[d.Family] = d.Help
		}
		for name := range d.FixedLabels {
			labels[name] = true
		}
		if d.ValueLabel != "" {
			labels[d.ValueLabel] = true
		}
		if d.Transform == "string_info" {
			labels["method"] = true
		}
	}
	descriptors := make(map[string]familyDescriptor, len(schemas))
	for family, labelSet := range schemas {
		names := make([]string, 0, len(labelSet))
		for name := range labelSet {
			names = append(names, name)
		}
		sort.Strings(names)
		descriptors[family] = familyDescriptor{prometheus.NewDesc(family, helps[family], names, nil), names}
	}
	return &Module{descriptors: descriptors,
		decodeErrors: prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: "mongodb", Subsystem: "server_status", Name: "decode_errors_total", Help: "Total modern server-status metric-family decode failures."}, []string{"family"}),
		unknownKeys:  prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: "mongodb", Subsystem: "server_status", Name: "unknown_keys_total", Help: "Total ignored unknown keys in bounded modern server-status families."}, []string{"family"})}
}
func (m *Module) Describe(ch chan<- *prometheus.Desc) {
	families := make([]string, 0, len(m.descriptors))
	for family := range m.descriptors {
		families = append(families, family)
	}
	sort.Strings(families)
	for _, family := range families {
		ch <- m.descriptors[family].desc
	}
	m.decodeErrors.Describe(ch)
	m.unknownKeys.Describe(ch)
}
func (m *Module) Collect(raw bson.Raw, ch chan<- prometheus.Metric) Report {
	grouped := make(map[string][]sample)
	failures := make(map[string]FamilyError)
	snapshots := snapshotValues(raw)
	for _, d := range metricDefinitions {
		value := lookupDefinition(raw, d)
		if value.Type == 0 || value.Type == bsontype.Null {
			continue
		}
		labels := copyLabels(d.FixedLabels)
		numeric, err := metricValue(value, d, labels, snapshots)
		if err != nil {
			if _, ok := failures[d.Family]; !ok {
				failures[d.Family] = FamilyError{d.Family, d.SourcePath, err}
			}
			continue
		}
		grouped[d.Family] = append(grouped[d.Family], sample{d, numeric, labels})
	}
	report := Report{}
	families := make([]string, 0, len(grouped))
	for family := range grouped {
		families = append(families, family)
	}
	sort.Strings(families)
	for _, family := range families {
		if _, failed := failures[family]; failed {
			continue
		}
		descriptor := m.descriptors[family]
		for _, candidate := range grouped[family] {
			values := make([]string, len(descriptor.labelNames))
			for i, name := range descriptor.labelNames {
				values[i] = candidate.labels[name]
			}
			kind := prometheus.GaugeValue
			if candidate.definition.Type == "counter" {
				kind = prometheus.CounterValue
			}
			ch <- prometheus.MustNewConstMetric(descriptor.desc, kind, candidate.value, values...)
			report.Emitted++
		}
	}
	failed := make([]string, 0, len(failures))
	for family := range failures {
		failed = append(failed, family)
	}
	sort.Strings(failed)
	for _, family := range failed {
		failure := failures[family]
		report.Errors = append(report.Errors, failure)
		m.decodeErrors.WithLabelValues(family).Inc()
	}
	m.recordUnknownKeys(raw)
	m.decodeErrors.Collect(ch)
	m.unknownKeys.Collect(ch)
	return report
}
func lookupDefinition(raw bson.Raw, d metricDefinition) bson.RawValue {
	segments := d.PathSegments
	if d.Transform == "timestamp_seconds" || d.Transform == "timestamp_increment" {
		segments = segments[:len(segments)-1]
	}
	return raw.Lookup(segments...)
}
func metricValue(value bson.RawValue, d metricDefinition, labels map[string]string, snapshots map[string]snapshotValue) (float64, error) {
	if d.Transform == "snapshot_seconds" || d.Transform == "snapshot_increment" {
		v := snapshots[d.SourcePath]
		if d.Transform == "snapshot_increment" {
			return v.increment, nil
		}
		return v.seconds, nil
	}
	if d.Transform == "string_info" {
		if value.Type != bsontype.String {
			return 0, fmt.Errorf("got %s, want string", value.Type)
		}
		labels["method"] = value.StringValue()
		return 1, nil
	}
	if d.ValueLabel != "" {
		if value.Type != bsontype.String {
			return 0, fmt.Errorf("got %s, want string", value.Type)
		}
		actual := value.StringValue()
		allowed := false
		for _, candidate := range d.AllowedValues {
			if actual == candidate {
				allowed = true
				break
			}
		}
		if !allowed {
			return 0, fmt.Errorf("value %q is outside bounded vocabulary", actual)
		}
		labels[d.ValueLabel] = actual
		return 1, nil
	}
	var numeric float64
	switch d.Transform {
	case "date_milliseconds":
		if value.Type != bsontype.DateTime {
			return 0, fmt.Errorf("got %s, want date", value.Type)
		}
		numeric = float64(value.DateTime())
	case "timestamp_seconds", "timestamp_increment":
		if value.Type != bsontype.Timestamp {
			return 0, fmt.Errorf("got %s, want timestamp", value.Type)
		}
		seconds, increment := value.Timestamp()
		if d.Transform == "timestamp_seconds" {
			numeric = float64(seconds)
		} else {
			numeric = float64(increment)
		}
	default:
		switch value.Type {
		case bsontype.Int32:
			numeric = float64(value.Int32())
		case bsontype.Int64:
			numeric = float64(value.Int64())
		case bsontype.Double:
			numeric = value.Double()
		case bsontype.Boolean:
			if value.Boolean() {
				numeric = 1
			}
		case bsontype.DateTime:
			numeric = float64(value.DateTime()) / 1000
		case bsontype.Timestamp:
			seconds, _ := value.Timestamp()
			numeric = float64(seconds)
		default:
			return 0, fmt.Errorf("got %s, want numeric, boolean, date, or timestamp", value.Type)
		}
	}
	return numeric * d.Conversion, nil
}
func copyLabels(source map[string]string) map[string]string {
	result := make(map[string]string, len(source)+1)
	for k, v := range source {
		result[k] = v
	}
	return result
}

func (m *Module) recordUnknownKeys(raw bson.Raw) {
	for _, family := range generatedDynamicFamilies {
		value := raw.Lookup(family.Root...)
		if value.Type != bsontype.EmbeddedDocument {
			continue
		}
		walkDynamic(value.Document(), nil, func(path string) {
			if _, ok := family.Allowed[path]; !ok {
				m.unknownKeys.WithLabelValues(family.Name).Inc()
			}
		})
	}
}
func walkDynamic(doc bson.Raw, prefix []string, visit func(string)) {
	elements, err := doc.Elements()
	if err != nil {
		return
	}
	for _, element := range elements {
		path := append(prefix, element.Key())
		value := element.Value()
		if value.Type == bsontype.EmbeddedDocument {
			walkDynamic(value.Document(), path, visit)
			continue
		}
		visit(strings.Join(path, "\x00"))
	}
}

var snapshotPattern = regexp.MustCompile("^([A-Z][a-z]{2}) ([ 0-9][0-9]) ([0-9]{2}):([0-9]{2}):([0-9]{2}):([0-9]+)$")

type snapshotParts struct {
	month                     time.Month
	day, hour, minute, second int
	increment                 uint32
}
type snapshotAnchor struct {
	year          int
	offsetSeconds int
	latestUnix    int64
}

func parseSnapshotParts(text string) (snapshotParts, error) {
	m := snapshotPattern.FindStringSubmatch(text)
	if len(m) != 7 {
		return snapshotParts{}, fmt.Errorf("invalid snapshot timestamp %q", text)
	}
	parsed, err := time.Parse("Jan 02 15:04:05", fmt.Sprintf("%s %s %s:%s:%s", m[1], m[2], m[3], m[4], m[5]))
	if err != nil {
		return snapshotParts{}, err
	}
	increment64, err := strconv.ParseUint(m[6], 10, 32)
	if err != nil {
		return snapshotParts{}, err
	}
	return snapshotParts{parsed.Month(), parsed.Day(), parsed.Hour(), parsed.Minute(), parsed.Second(), uint32(increment64)}, nil
}
func anchorLatestSnapshot(parts snapshotParts, localTime time.Time, logicalSeconds uint32) (snapshotAnchor, bool) {
	if logicalSeconds == 0 {
		return snapshotAnchor{}, false
	}
	target := int64(logicalSeconds)
	bestScore := int64(math.MaxInt64)
	best := snapshotAnchor{}
	ties := 0
	for year := localTime.Year() - 1; year <= localTime.Year()+1; year++ {
		for offsetMinutes := -14 * 60; offsetMinutes <= 14*60; offsetMinutes += 15 {
			location := time.FixedZone("snapshot", offsetMinutes*60)
			candidate := time.Date(year, parts.month, parts.day, parts.hour, parts.minute, parts.second, 0, location).Unix()
			score := candidate - target
			if score < 0 {
				score = -score
			}
			if score < bestScore {
				bestScore = score
				best = snapshotAnchor{year, offsetMinutes * 60, candidate}
				ties = 1
			} else if score == bestScore {
				ties++
			}
		}
	}
	return best, ties == 1 && bestScore <= 48*60*60
}
func anchoredSnapshot(parts snapshotParts, anchor snapshotAnchor, latest bool) (snapshotValue, bool) {
	location := time.FixedZone("snapshot", anchor.offsetSeconds)
	candidate := time.Date(anchor.year, parts.month, parts.day, parts.hour, parts.minute, parts.second, 0, location).Unix()
	if !latest && candidate > anchor.latestUnix {
		candidate = time.Date(anchor.year-1, parts.month, parts.day, parts.hour, parts.minute, parts.second, 0, location).Unix()
	}
	if candidate > anchor.latestUnix || anchor.latestUnix-candidate > 366*24*60*60 {
		return snapshotValue{}, false
	}
	return snapshotValue{float64(candidate), float64(parts.increment)}, true
}
func snapshotValues(raw bson.Raw) map[string]snapshotValue {
	result := map[string]snapshotValue{}
	local := raw.Lookup("localTime")
	logical := raw.Lookup("$clusterTime", "clusterTime")
	if local.Type != bsontype.DateTime || logical.Type != bsontype.Timestamp {
		return result
	}
	logicalSeconds, _ := logical.Timestamp()
	latestPath := "wiredTiger['snapshot-window-settings']['latest majority snapshot timestamp available']"
	oldestPath := "wiredTiger['snapshot-window-settings']['oldest majority snapshot timestamp available']"
	latest := raw.Lookup("wiredTiger", "snapshot-window-settings", "latest majority snapshot timestamp available")
	if latest.Type != bsontype.String {
		return result
	}
	parts, err := parseSnapshotParts(latest.StringValue())
	if err != nil {
		return result
	}
	anchor, ok := anchorLatestSnapshot(parts, time.Unix(0, local.DateTime()*int64(time.Millisecond)), logicalSeconds)
	if !ok {
		return result
	}
	value, ok := anchoredSnapshot(parts, anchor, true)
	if !ok {
		return result
	}
	result[latestPath] = value
	oldest := raw.Lookup("wiredTiger", "snapshot-window-settings", "oldest majority snapshot timestamp available")
	if oldest.Type != bsontype.String {
		return result
	}
	oldestParts, err := parseSnapshotParts(oldest.StringValue())
	if err != nil {
		return result
	}
	oldestValue, ok := anchoredSnapshot(oldestParts, anchor, false)
	if ok {
		result[oldestPath] = oldestValue
	}
	return result
}
