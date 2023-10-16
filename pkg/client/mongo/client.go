package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"rest_api/config"
	"time"
)

func GetMongoEngine() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	conf := config.New()
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s/?retryWrites=true&w=majority", conf.MongoDB.MongoUser, conf.MongoDB.MongoPassword, conf.MongoDB.MongoDomain)
	fmt.Println(connectionString)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}
	return client, nil
}
