package mongo_client

import (
	"context"
	"time"

	"github.com/vladislavkovaliov/ledger/internal/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewMongoClient(cfg *config.Config) (*mongo.Client, *mongo.Database, error) {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(cfg.DatabaseUrl))

	if err != nil {
		return nil, nil, err
	}

	db := client.Database("ledger")

	return client, db, nil
}
