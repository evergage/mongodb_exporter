package serverstatusv5

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Contract struct {
	Version                      int                 `json:"version"`
	MongoDBVersions              []string            `json:"mongodb_versions"`
	Classifications              []string            `json:"classifications"`
	SeriesCeilings               SeriesCeilings      `json:"series_ceilings"`
	ScrapeDurationSecondsCeiling float64             `json:"scrape_duration_seconds_ceiling"`
	DynamicVocabularies          map[string][]string `json:"dynamic_vocabularies"`
	Entries                      []ContractEntry     `json:"entries"`
}
type SeriesCeilings struct {
	PerFamily int `json:"per_family"`
	Primary   int `json:"primary"`
	Secondary int `json:"secondary"`
	Union     int `json:"union"`
}
type CoverageLeaf struct {
	Path         string
	PathSegments []string
	Roles        []string
}

type ContractEntry struct {
	Path              string           `json:"path"`
	PathSegments      []string         `json:"path_segments"`
	Roles             []string         `json:"roles"`
	Classification    string           `json:"classification"`
	Reason            string           `json:"reason"`
	DecisionSource    string           `json:"decision_source"`
	LegacyFamily      string           `json:"legacy_family,omitempty"`
	Reviewed          bool             `json:"reviewed,omitempty"`
	ConsumedBy        string           `json:"consumed_by,omitempty"`
	Metric            *MetricContract  `json:"metric,omitempty"`
	AdditionalMetrics []MetricContract `json:"additional_metrics,omitempty"`
}
type MetricContract struct {
	Family        string            `json:"family"`
	Type          string            `json:"type"`
	Unit          string            `json:"unit"`
	Conversion    float64           `json:"conversion"`
	Help          string            `json:"help"`
	Labels        map[string]string `json:"labels"`
	Transform     string            `json:"transform,omitempty"`
	ValueLabel    string            `json:"value_label,omitempty"`
	AllowedValues []string          `json:"allowed_values"`
	SeriesCeiling int               `json:"series_ceiling"`
}

func ParseContract(data []byte) (*Contract, error) {
	var c Contract
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}
	return &c, nil
}

func (c *Contract) Audit(leaves []CoverageLeaf) error {
	valid := map[string]bool{}
	for _, value := range c.Classifications {
		valid[value] = true
	}
	fixture := map[string]CoverageLeaf{}
	fixtureCounts := map[string]int{}
	for _, leaf := range leaves {
		fixture[leaf.Path] = leaf
		fixtureCounts[leaf.Path]++
	}
	entryCounts := map[string]int{}
	entryByPath := map[string]ContractEntry{}
	var invalid []string
	for _, entry := range c.Entries {
		entryCounts[entry.Path]++
		entryByPath[entry.Path] = entry
		if !valid[entry.Classification] || entry.Reason == "" || entry.DecisionSource == "" || len(entry.PathSegments) == 0 {
			invalid = append(invalid, entry.Path)
		}
		if leaf, ok := fixture[entry.Path]; ok {
			if !sameStrings(entry.PathSegments, leaf.PathSegments) || !sameStrings(entry.Roles, leaf.Roles) {
				invalid = append(invalid, entry.Path)
			}
		}
		switch entry.Classification {
		case "legacy":
			if entry.LegacyFamily == "" || entry.Metric != nil {
				invalid = append(invalid, entry.Path)
			}
		case "modern":
			if entry.Metric == nil && entry.ConsumedBy == "" {
				invalid = append(invalid, entry.Path)
			}
		case "drop":
			if !entry.Reviewed || entry.Metric != nil {
				invalid = append(invalid, entry.Path)
			}
		}
	}
	for _, entry := range c.Entries {
		if entry.ConsumedBy == "" {
			continue
		}
		target, ok := entryByPath[entry.ConsumedBy]
		if !ok || target.Metric == nil || entry.ConsumedBy == entry.Path {
			invalid = append(invalid, entry.Path)
		}
	}
	var missing, duplicate, stale []string
	for path, count := range fixtureCounts {
		if count > 1 || entryCounts[path] > 1 {
			duplicate = append(duplicate, path)
		}
		if entryCounts[path] == 0 {
			missing = append(missing, path)
		}
	}
	for path, count := range entryCounts {
		if count > 1 {
			duplicate = append(duplicate, path)
		}
		if fixtureCounts[path] == 0 {
			stale = append(stale, path)
		}
	}
	sort.Strings(missing)
	sort.Strings(duplicate)
	sort.Strings(stale)
	sort.Strings(invalid)
	if len(missing)+len(duplicate)+len(stale)+len(invalid) > 0 {
		return fmt.Errorf("coverage audit failed: missing=%v duplicate=%v stale=%v invalid=%v", missing, duplicate, stale, invalid)
	}
	return nil
}

func sameStrings(left, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}

func (c *Contract) ValidatePolicy() error {
	if c.ScrapeDurationSecondsCeiling <= 0 {
		return fmt.Errorf("scrape duration ceiling must be positive")
	}
	familyCounts := map[string]int{}
	roleCounts := map[string]int{"primary": 0, "secondary": 0}
	for _, e := range c.Entries {
		metrics := make([]*MetricContract, 0, 1+len(e.AdditionalMetrics))
		if e.Metric != nil {
			metrics = append(metrics, e.Metric)
		}
		for i := range e.AdditionalMetrics {
			metrics = append(metrics, &e.AdditionalMetrics[i])
		}
		for _, m := range metrics {
			if m.Family == "" || m.Type == "" || m.Unit == "" || m.Help == "" || m.Conversion == 0 {
				return fmt.Errorf("%s has an incomplete metric mapping", e.Path)
			}
			if e.DecisionSource != "existing modern contract" && m.Type == "counter" && !strings.HasSuffix(m.Family, "_total") {
				return fmt.Errorf("counter %s does not end in _total", m.Family)
			}
			if e.DecisionSource != "existing modern contract" && m.Conversion != 1 {
				return fmt.Errorf("new mapping %s conversion=%v, want 1", e.Path, m.Conversion)
			}
			for name, value := range m.Labels {
				if strings.HasPrefix(name, "lower_bound_") {
					if _, err := strconv.ParseFloat(value, 64); err != nil {
						return fmt.Errorf("%s has nonnumeric %s=%q", e.Path, name, value)
					}
				}
			}
			familyCounts[m.Family]++
			for _, role := range e.Roles {
				roleCounts[role]++
			}
		}
	}
	maxFamily, union := 0, 0
	for family, count := range familyCounts {
		union += count
		if count > maxFamily {
			maxFamily = count
		}
		for _, e := range c.Entries {
			if e.Metric != nil && e.Metric.Family == family && e.Metric.SeriesCeiling != count {
				return fmt.Errorf("%s ceiling is %d, want %d", family, e.Metric.SeriesCeiling, count)
			}
			for _, m := range e.AdditionalMetrics {
				if m.Family == family && m.SeriesCeiling != count {
					return fmt.Errorf("%s ceiling is %d, want %d", family, m.SeriesCeiling, count)
				}
			}
		}
	}
	got := SeriesCeilings{maxFamily, roleCounts["primary"], roleCounts["secondary"], union}
	if got != c.SeriesCeilings {
		return fmt.Errorf("derived ceilings %+v do not match %+v", got, c.SeriesCeilings)
	}
	wantSizes := map[string]int{"commands": 60, "aggregation_stages": 49, "operators_expressions": 141, "operators_groupAccumulators": 14, "operators_match": 35, "operators_windowAccumulators": 20}
	for name, want := range wantSizes {
		values := c.DynamicVocabularies[name]
		if len(values) != want {
			return fmt.Errorf("%s vocabulary has %d values, want %d", name, len(values), want)
		}
		seen := map[string]bool{}
		for _, v := range values {
			if seen[v] {
				return fmt.Errorf("%s vocabulary repeats %q", name, v)
			}
			seen[v] = true
		}
	}
	return nil
}

var wiredTigerRangePattern = regexp.MustCompile(`(?:histogram - |\) - )(\d+)(?:-|(?:ms|us)\+| and higher)`)

func wiredTigerHistogramLowerBound(description string) (string, error) {
	m := wiredTigerRangePattern.FindStringSubmatch(description)
	if len(m) != 2 {
		return "", fmt.Errorf("ambiguous WiredTiger histogram range %q", description)
	}
	return m[1], nil
}
