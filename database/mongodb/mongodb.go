package mongodb

import (
	"context"
	"errors"
	"fmt"

	"github.com/lilac-project/backend-engine/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	DBRead, DBWrite *mongo.Database
	Ctx             *context.Context
}

func InitMongoRepository(ctx *context.Context, configs config.Configs) (*MongoRepository, error) {
	if configs.Mongo == nil {
		return nil, errors.New("failed to connect to DB, msg: configs is not set")
	}
	dbRead, err := DBRead(ctx, *configs.Mongo)
	if err != nil {
		return nil, err
	}
	dbWrite, err := DBRead(ctx, *configs.Mongo)
	if err != nil {
		return nil, err
	}
	return &MongoRepository{
		DBRead:  dbRead,
		DBWrite: dbWrite,
		Ctx:     ctx,
	}, nil
}

func DBRead(ctx *context.Context, mongoConfig config.MongoConfig) (*mongo.Database, error) {
	username := mongoConfig.User
	password := mongoConfig.Password
	host := mongoConfig.Host
	db := mongoConfig.DB
	url := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority",
		username, password, host, db)
	client, err := Connect(url, *ctx)

	if err != nil {
		return nil, err
	}
	return client.Database(db), nil
}

func DBWrite(ctx *context.Context, mongoConfig config.MongoConfig) (*mongo.Database, error) {
	username := mongoConfig.User
	password := mongoConfig.Password
	host := mongoConfig.Host
	db := mongoConfig.DB
	url := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority",
		username, password, host, db)
	client, err := Connect(url, *ctx)

	if err != nil {
		return nil, err
	}
	return client.Database(db), nil
}

func Connect(url string, ctx context.Context) (*mongo.Client, error) {
	clientOptions := options.Client()
	clientOptions.SetMaxPoolSize(100)
	clientOptions.SetMinPoolSize(10)

	client, err := mongo.NewClient(clientOptions.ApplyURI(url))

	if err != nil {
		return nil, err
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func CloseConnection(client *mongo.Client, ctx *context.Context) {
	client.Disconnect(*ctx)
}
