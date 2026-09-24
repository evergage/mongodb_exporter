package serverstatusv5

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

func TestFocusedCanonicalExtendedJSONTypes(t *testing.T) {
	fixture := []byte(`{
      "int32":{"$numberInt":"7"},
      "int64":{"$numberLong":"9007199254740993"},
      "double":{"$numberDouble":"1.5"},
      "date":{"$date":{"$numberLong":"1660000000000"}},
      "timestamp":{"$timestamp":{"t":42,"i":3}},
      "boolean":false,
      "null":null,
      "objectID":{"$oid":"000000000000000000000001"},
      "binary":{"$binary":{"base64":"AAE=","subType":"00"}}
    }`)
	var raw bson.Raw
	if err := bson.UnmarshalExtJSON(fixture, true, &raw); err != nil {
		t.Fatal(err)
	}
	expected := map[string]bsontype.Type{
		"int32": bsontype.Int32, "int64": bsontype.Int64, "double": bsontype.Double,
		"date": bsontype.DateTime, "timestamp": bsontype.Timestamp, "boolean": bsontype.Boolean,
		"null": bsontype.Null, "objectID": bsontype.ObjectID, "binary": bsontype.Binary,
	}
	for key, want := range expected {
		if got := raw.Lookup(key).Type; got != want {
			t.Errorf("%s type = %s, want %s", key, got, want)
		}
	}
}

func TestFixtureLogicalTimesUseWireTimestampType(t *testing.T) {
	raw := loadFixtureRaw(t, "primary")
	paths := [][]string{
		{"$clusterTime", "clusterTime"},
		{"$configServerState", "opTime", "ts"},
		{"$gleStats", "lastOpTime"},
		{"lastCommittedOpTime"},
		{"operationTime"},
		{"repl", "lastWrite", "majorityOpTime", "ts"},
		{"repl", "lastWrite", "opTime", "ts"},
		{"sharding", "lastSeenConfigServerOpTime", "ts"},
		{"storageEngine", "oldestRequiredTimestampForCrashRecovery"},
		{"wiredTiger", "oplog", "visibility timestamp"},
		{"wiredTiger", "snapshot-window-settings", "min pinned timestamp"},
	}
	for _, path := range paths {
		if got := raw.Lookup(path...).Type; got != bsontype.Timestamp {
			t.Errorf("%v type = %s, want %s", path, got, bsontype.Timestamp)
		}
	}
}
