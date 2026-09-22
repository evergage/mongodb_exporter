package mongod

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/percona/mongodb_exporter/testutils"
)

func Test_ParserTopStatus(t *testing.T) {
	raw := &TopStatusRaw{}
	client := testutils.MustGetConnectedMongodClient(context.TODO(), t)
	defer client.Disconnect(context.TODO())
	err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"top", 1}}).Decode(&raw)
	if err != nil {
		t.Fatal(err)
	}

	topStatus := raw.TopStatus()
	assert.NotEmpty(t, topStatus.TopStats)
	positive := false
	for _, stats := range topStatus.TopStats {
		if stats.Total.Time > 0 && stats.Total.Count > 0 {
			positive = true
			break
		}
	}
	assert.True(t, positive, "expected at least one observed namespace")
}
