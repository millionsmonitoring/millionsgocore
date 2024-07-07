package initializers

import (
	"context"
	"log/slog"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// default mongo db provider
func InitMongoDB(ctx context.Context) (*mongo.Client, error) {
	mongoUrl := os.Getenv("MONGODB_URL")
	if mongoUrl == "" {
		slog.Error("MONGODB_URL is not set in the environment")
		panic("MONGODB_URL is not set in the environment")
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		return nil, err
	}
	// Check the connection
	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return client, nil
}

func CloseMongoDB(ctx context.Context, c *mongo.Client) error {
	return c.Disconnect(ctx)
}
