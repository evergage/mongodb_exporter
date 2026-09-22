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

func (c *Contract) Audit(paths []string) error {
	valid := map[string]bool{}
	for _, v := range c.Classifications {
		valid[v] = true
	}
	fixture := map[string]int{}
	for _, p := range paths {
		fixture[p]++
	}
	entries := map[string]int{}
	var invalid []string
	for _, e := range c.Entries {
		entries[e.Path]++
		if !valid[e.Classification] || e.Reason == "" || e.DecisionSource == "" || len(e.PathSegments) == 0 {
			invalid = append(invalid, e.Path)
		}
		switch e.Classification {
		case "legacy":
			if e.LegacyFamily == "" || e.Metric != nil {
				invalid = append(invalid, e.Path)
			}
		case "modern":
			if e.Metric == nil && e.ConsumedBy == "" {
				invalid = append(invalid, e.Path)
			}
		case "drop":
			if !e.Reviewed || e.Metric != nil {
				invalid = append(invalid, e.Path)
			}
		}
	}
	var missing, duplicate, stale []string
	for p, n := range fixture {
		if n > 1 || entries[p] > 1 {
			duplicate = append(duplicate, p)
		}
		if entries[p] == 0 {
			missing = append(missing, p)
		}
	}
	for p, n := range entries {
		if n > 1 {
			duplicate = append(duplicate, p)
		}
		if fixture[p] == 0 {
			stale = append(stale, p)
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
