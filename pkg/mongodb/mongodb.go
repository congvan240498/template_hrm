package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"hrm/config"
	"hrm/pkg/logger"
)

type DatabaseStorage struct {
	db     *mongo.Database
	client *mongo.Client
}

var dbStorage *DatabaseStorage

func ConnectMongoDB(ctx context.Context, config *config.MongoDBConfig) (*DatabaseStorage, error) {
	if dbStorage != nil {
		return dbStorage, nil
	}

	client, db, err := connect(ctx, config)
	if err != nil {
		return nil, err
	}

	dbStorage = &DatabaseStorage{
		db:     db,
		client: client,
	}
	return dbStorage, nil

}

func connect(ctx context.Context, config *config.MongoDBConfig) (*mongo.Client, *mongo.Database, error) {
	log := logger.GetLogger()

	ctxNew, cc := context.WithTimeout(ctx, 30*time.Second)
	defer cc()

	clientOpts := options.Client().ApplyURI(config.DatabaseURI)
	clientOpts.SetMaxPoolSize(100)

	client, err := mongo.Connect(ctxNew, clientOpts)
	if err != nil {
		log.Error().Msg("connect mongo failed")
		return nil, nil, err
	}

	if err = client.Ping(ctxNew, readpref.Primary()); err != nil {
		log.Error().Msg("ping mongo failed")
		return nil, nil, err
	}

	log.Info().Msgf("connect mongodb successfully: db_name=%s", config.DatabaseName)
	return client, client.Database(config.DatabaseName), nil
}
