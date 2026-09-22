package testutils

import (
	"context"
	"os"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongodURI() string {
	if value := os.Getenv("TEST_MONGODB_URI"); value != "" {
		return value
	}
	return "mongodb://127.0.0.1:27017/admin"
}
func ReplSetURI() string {
	if value := os.Getenv("TEST_REPLSET_URI"); value != "" {
		return value
	}
	return "mongodb://127.0.0.1:27019/admin"
}
func MongosURI() string {
	if value := os.Getenv("TEST_MONGOS_URI"); value != "" {
		return value
	}
	return "mongodb://127.0.0.1:27017/admin"
}

// MustGetConnectedReplSetClient return mongo.Client instance connected to server started in replicaSet mode.
func MustGetConnectedReplSetClient(ctx context.Context, t *testing.T) *mongo.Client {
	opts := options.Client().
		ApplyURI(ReplSetURI()).
		SetDirect(true).SetServerSelectionTimeout(time.Second)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		t.Fatalf("Couldn't connect to MongoDB instance, reason: %v", err)
	}

	return client
}

// MustGetConnectedMongodClient return mongo.Client instance connected to server started in single mode.
func MustGetConnectedMongodClient(ctx context.Context, t *testing.T) *mongo.Client {
	opts := options.Client().
		ApplyURI(MongodURI()).
		SetDirect(true).SetServerSelectionTimeout(time.Second)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		t.Fatalf("Couldn't connect to MongoDB instance, reason: %v", err)
	}

	return client
}
