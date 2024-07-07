// Package db provides helper functions that make interfacing with the MongoDB Go driver library easier
package mongoboiler

import (
	"context"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	db     *mongo.Database
	client *mongo.Client
}

func New(client *mongo.Client, name string) *DB {
	return &DB{client.Database(name), client}
}

func (db DB) Disconnect(ctx context.Context) error {
	return db.client.Disconnect(ctx)
}

// Collection is the wrapper for Mongo Collection
type Collection struct {
	collection *mongo.Collection
}

func (wrapper *DB) NewCollection(collectionName string) *Collection {
	return &Collection{wrapper.db.Collection(collectionName)}
}

// Drop drops the current Collection (collection)
func (c Collection) Drop(ctx context.Context) error {
	return c.collection.Drop(ctx)
}

// FindOne finds first document that satisfies filter and fills res with the un marshaled document.
func (c Collection) FindOne(ctx context.Context, filter bson.D, res any) error {
	err := c.collection.FindOne(ctx, filter).Decode(res)
	if err != nil {
		return err
	}
	return nil
}

// FindMany iterates cursor of all docs matching filter and fills res with un marshalled documents.
func (c Collection) FindMany(ctx context.Context, filter bson.D, res *[]any) error {
	arrType := reflect.TypeOf(res).Elem()
	cursor, err := c.collection.Find(ctx, filter)

	for cursor.Next(ctx) {
		doc := reflect.New(arrType).Interface()
		err := cursor.Decode(&doc)
		if err != nil {
			return err
		}
		*res = append(*res, doc)
	}

	// un marshall fail
	if cursor.Err() != nil {
		return err
	}

	// Close cursor after we're done with it
	cursor.Close(ctx)
	return nil
}

// UpdateOne updates single document matching filter and applies update to it.
// Returns number of documents matched and modified. Should always be either 0 or 1.
func (c Collection) UpdateOne(ctx context.Context, filter, update bson.D) (int64, int64, error) {
	updateRes, err := c.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, 0, err
	}
	return updateRes.MatchedCount, updateRes.ModifiedCount, nil
}

// UpdateMany updates all documents matching the filter by applying the update query on it.
// Returns number of documents matched and modified.
func (c Collection) UpdateMany(ctx context.Context, filter, update bson.D) (int64, int64, error) {
	updateRes, err := c.collection.UpdateMany(ctx, filter, update)
	if err != nil {
		return 0, 0, err
	}
	return updateRes.MatchedCount, updateRes.ModifiedCount, nil
}

// InsertOne inserts a single struct as a document into the database and returns its ID.
// Returns inserted ID
func (c Collection) InsertOne(ctx context.Context, new any) (any, error) {
	insertRes, err := c.collection.InsertOne(ctx, new)
	if err != nil {
		return "", err
	}
	return insertRes.InsertedID, nil
}

// InsertMany takes a slice of structs, inserts them into the database.
// Returns list of inserted IDs
func (c Collection) InsertMany(ctx context.Context, new []any) (any, error) {
	insertRes, err := c.collection.InsertMany(ctx, new)
	if err != nil {
		return "", err
	}
	return insertRes.InsertedIDs, nil
}

// DeleteOne deletes single document that match the bson.D filter
func (c Collection) DeleteOne(ctx context.Context, filter bson.D) error {
	_, err := c.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMany deletes all documents that match the bson.D filter
func (c Collection) DeleteMany(ctx context.Context, filter bson.D) error {
	_, err := c.collection.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
