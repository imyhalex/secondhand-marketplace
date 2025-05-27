package utils

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// reference: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo

func NewMongo(ctx context.Context, uri String) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cacel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return client, client.Ping(ctx, nil)
}
