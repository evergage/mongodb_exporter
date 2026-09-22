package serverstatusv5

import (
	"encoding/json"
	"fmt"
	"sort"
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
	Path           string          `json:"path"`
	Roles          []string        `json:"roles"`
	Classification string          `json:"classification"`
	Reason         string          `json:"reason,omitempty"`
	Metric         *MetricContract `json:"metric,omitempty"`
}
type MetricContract struct {
	Family        string            `json:"family"`
	Type          string            `json:"type"`
	Unit          string            `json:"unit"`
	Conversion    float64           `json:"conversion"`
	Help          string            `json:"help"`
	Labels        map[string]string `json:"labels"`
	ValueLabel    string            `json:"value_label"`
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
	valid := make(map[string]bool, len(c.Classifications))
	for _, v := range c.Classifications {
		valid[v] = true
	}
	seen := make(map[string]int, len(c.Entries))
	for _, e := range c.Entries {
		seen[e.Path]++
		if !valid[e.Classification] {
			return fmt.Errorf("unknown classification %q for %s", e.Classification, e.Path)
		}
		if e.Classification == "new_metric" && e.Metric == nil {
			return fmt.Errorf("new metric %s has no metric contract", e.Path)
		}
	}
	fixture := make(map[string]bool, len(paths))
	for _, p := range paths {
		fixture[p] = true
	}
	var missing, duplicate, stale []string
	for p := range fixture {
		if seen[p] == 0 {
			missing = append(missing, p)
		}
	}
	for p, n := range seen {
		if n > 1 {
			duplicate = append(duplicate, p)
		}
		if !fixture[p] {
			stale = append(stale, p)
		}
	}
	sort.Strings(missing)
	sort.Strings(duplicate)
	sort.Strings(stale)
	if len(missing)+len(duplicate)+len(stale) > 0 {
		return fmt.Errorf("coverage audit failed: missing=%v duplicate=%v stale=%v", missing, duplicate, stale)
	}
	return nil
}
func (c *Contract) ValidatePolicy() error {
	if c.ScrapeDurationSecondsCeiling <= 0 {
		return fmt.Errorf("scrape duration ceiling must be positive")
	}
	familyCounts := make(map[string]int)
	roleCounts := map[string]int{"primary": 0, "secondary": 0}
	for _, entry := range c.Entries {
		if entry.Classification != "new_metric" {
			if entry.Reason == "" {
				return fmt.Errorf("%s has no omission rationale", entry.Path)
			}
			continue
		}
		metric := entry.Metric
		if metric == nil || metric.Family == "" || metric.Type == "" || metric.Unit == "" || metric.Help == "" || metric.Conversion == 0 {
			return fmt.Errorf("%s has an incomplete metric mapping", entry.Path)
		}
		possible := 1
		if metric.ValueLabel != "" {
			if len(metric.AllowedValues) == 0 {
				return fmt.Errorf("%s has an unbounded value label", entry.Path)
			}
			possible = len(metric.AllowedValues)
		}
		familyCounts[metric.Family] += possible
		for _, role := range entry.Roles {
			roleCounts[role] += possible
		}
	}
	maxFamily, union := 0, 0
	for family, count := range familyCounts {
		union += count
		if count > maxFamily {
			maxFamily = count
		}
		for _, entry := range c.Entries {
			if entry.Metric != nil && entry.Metric.Family == family && entry.Metric.SeriesCeiling != count {
				return fmt.Errorf("%s ceiling is %d, want %d", family, entry.Metric.SeriesCeiling, count)
			}
		}
	}
	if maxFamily != c.SeriesCeilings.PerFamily || roleCounts["primary"] != c.SeriesCeilings.Primary || roleCounts["secondary"] != c.SeriesCeilings.Secondary || union != c.SeriesCeilings.Union {
		return fmt.Errorf("derived ceilings family=%d primary=%d secondary=%d union=%d do not match %+v", maxFamily, roleCounts["primary"], roleCounts["secondary"], union, c.SeriesCeilings)
	}
	for name, values := range c.DynamicVocabularies {
		if len(values) == 0 {
			return fmt.Errorf("%s vocabulary is empty", name)
		}
		seen := make(map[string]bool, len(values))
		for _, value := range values {
			if seen[value] {
				return fmt.Errorf("%s vocabulary repeats %s", name, value)
			}
			seen[value] = true
		}
	}
	return nil
}
