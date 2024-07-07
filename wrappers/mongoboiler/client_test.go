package mongoboiler_test

import (
	"context"
	"os"
	"testing"

	"github.com/millionsmonitoring/millionsgocore/wrappers/mongoboiler"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupTestDB(t *testing.T) (*mongo.Client, *mongoboiler.DB, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// testDB := client.Database("testdb")
	return client, mongoboiler.New(client, "testdb"), nil
}

func cleanupTestDB(t *testing.T, client *mongo.Client) {
	err := client.Disconnect(context.Background())
	if err != nil {
		t.Fatalf("Failed to disconnect from MongoDB: %v", err)
	}
}

func TestCollection_InsertOne(t *testing.T) {
	client, db, _ := setupTestDB(t)
	defer cleanupTestDB(t, client)

	coll := db.NewCollection("testcollection")

	// Create your test document here
	testDocument := map[string]any{
		"data": "test",
	}

	// Insert the test document
	insertedID, err := coll.InsertOne(context.Background(), testDocument)
	if err != nil {
		t.Fatalf("InsertOne failed: %v", err)
	}

	t.Logf("Inserted ID: %v", insertedID)
}

func TestCollection_FindOne(t *testing.T) {
	client, db, _ := setupTestDB(t)
	defer cleanupTestDB(t, client)

	coll := db.NewCollection("testcollection")

	testFilter := bson.D{{Key: "data", Value: "test"}}
	var result map[string]any

	err := coll.FindOne(context.Background(), testFilter, &result)
	if err != nil {
		t.Fatalf("FindOne failed: %v", err)
	}

	t.Logf("Found document: %+v", result)
}

func TestCollection_FindMany(t *testing.T) {
	client, db, _ := setupTestDB(t)
	defer cleanupTestDB(t, client)

	coll := db.NewCollection("testcollection")

	testFilter := bson.D{{Key: "data", Value: "test"}}
	var result []any

	err := coll.FindMany(context.Background(), testFilter, &result)
	if err != nil {
		t.Fatalf("FindOne failed: %v", err)
	}

	t.Logf("Found document: %+v", result)
}

func TestMain(m *testing.M) {
	// Setup code, if any
	retCode := m.Run()
	// Teardown code, if any
	os.Exit(retCode)
}
