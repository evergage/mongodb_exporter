package serverstatusv5

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/percona/mongodb_exporter/collector/mongod"
	"github.com/percona/mongodb_exporter/collector/serverstatusv5/fixturetool"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

var legacyFamilyPattern = regexp.MustCompile(`fqName: "([^"]+)"`)

func TestLegacyOwnershipClaimsAreObservable(t *testing.T) {
	contract := loadPolicyContract(t)
	sources := map[string][]byte{}
	for _, role := range []string{"primary", "secondary"} {
		data, err := ioutil.ReadFile("testdata/source/server-status-" + role + "-5.0.34.json")
		if err != nil {
			t.Fatal(err)
		}
		sources[role] = data
	}
	baseline := map[string]map[string][]string{}
	for role, source := range sources {
		raw, err := sourceToRaw(source)
		if err != nil {
			t.Fatal(err)
		}
		baseline[role], err = legacySamples(raw)
		if err != nil {
			t.Fatal(err)
		}
	}
	claims := 0
	for _, entry := range contract.Entries {
		if entry.Classification != "legacy" {
			continue
		}
		claims++
		for _, role := range entry.Roles {
			mutated, err := mutateSourceLeaf(sources[role], entry.PathSegments)
			if err != nil {
				t.Fatalf("%s (%s): %v", entry.Path, role, err)
			}
			raw, err := sourceToRaw(mutated)
			if err != nil {
				t.Fatalf("%s (%s): %v", entry.Path, role, err)
			}
			after, err := legacySamples(raw)
			if err != nil {
				t.Fatalf("%s (%s): %v", entry.Path, role, err)
			}
			changed := false
			for _, family := range strings.Split(entry.LegacyFamily, ",") {
				family = strings.TrimSpace(family)
				if !equalStrings(baseline[role][family], after[family]) {
					changed = true
					break
				}
			}
			if !changed {
				t.Errorf("%s (%s) did not change claimed legacy family %s", entry.Path, role, entry.LegacyFamily)
			}
		}
	}
	if claims != 144 {
		t.Fatalf("legacy claims=%d want 144", claims)
	}
}

func sourceToRaw(source []byte) (bson.Raw, error) {
	data, err := fixturetool.ReconstructJSON(source)
	if err != nil {
		return nil, err
	}
	var raw bson.Raw
	if err := bson.UnmarshalExtJSON(data, true, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

func mutateSourceLeaf(source []byte, segments []string) ([]byte, error) {
	decoder := json.NewDecoder(strings.NewReader(string(source)))
	decoder.UseNumber()
	var value interface{}
	if err := decoder.Decode(&value); err != nil {
		return nil, err
	}
	current := value
	for _, segment := range segments[:len(segments)-1] {
		switch container := current.(type) {
		case map[string]interface{}:
			current = container[segment]
		case []interface{}:
			index, err := strconv.Atoi(segment)
			if err != nil {
				return nil, err
			}
			current = container[index]
		default:
			return nil, fmt.Errorf("segment %q traverses %T", segment, current)
		}
	}
	last := segments[len(segments)-1]
	mutate := func(original interface{}) (interface{}, error) {
		switch typed := original.(type) {
		case json.Number:
			if strings.ContainsAny(typed.String(), ".eE") {
				number, err := strconv.ParseFloat(typed.String(), 64)
				if err != nil {
					return nil, err
				}
				return json.Number(strconv.FormatFloat(number+1024, 'g', -1, 64)), nil
			}
			number, err := strconv.ParseInt(typed.String(), 10, 64)
			if err != nil {
				return nil, err
			}
			return json.Number(strconv.FormatInt(number+1048576, 10)), nil
		case bool:
			return !typed, nil
		case string:
			if parsed, err := time.Parse(time.RFC3339Nano, typed); err == nil {
				return parsed.Add(time.Minute).Format(time.RFC3339Nano), nil
			}
			if len(typed) == 24 {
				replacement := byte('0')
				if typed[23] == replacement {
					replacement = '1'
				}
				return typed[:23] + string(replacement), nil
			}
			return typed + "-probe", nil
		default:
			return nil, fmt.Errorf("cannot mutate %T", original)
		}
	}
	switch container := current.(type) {
	case map[string]interface{}:
		next, err := mutate(container[last])
		if err != nil {
			return nil, err
		}
		container[last] = next
	case []interface{}:
		index, err := strconv.Atoi(last)
		if err != nil {
			return nil, err
		}
		next, err := mutate(container[index])
		if err != nil {
			return nil, err
		}
		container[index] = next
	default:
		return nil, fmt.Errorf("leaf parent is %T", current)
	}
	return json.Marshal(value)
}

func legacySamples(raw bson.Raw) (map[string][]string, error) {
	var status mongod.ServerStatus
	if err := bson.Unmarshal(raw, &status); err != nil {
		return nil, err
	}
	metrics := make(chan prometheus.Metric, 10000)
	status.Export(metrics)
	close(metrics)
	result := map[string][]string{}
	for metric := range metrics {
		match := legacyFamilyPattern.FindStringSubmatch(metric.Desc().String())
		if len(match) != 2 {
			return nil, fmt.Errorf("cannot parse descriptor %s", metric.Desc())
		}
		var wire dto.Metric
		if err := metric.Write(&wire); err != nil {
			return nil, err
		}
		result[match[1]] = append(result[match[1]], metric.Desc().String()+fmt.Sprintf("|%v", &wire))
	}
	for family := range result {
		sort.Strings(result[family])
	}
	return result, nil
}
func equalStrings(left, right []string) bool {
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
