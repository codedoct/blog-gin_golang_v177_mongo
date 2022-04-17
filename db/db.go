package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"

	"blog-gin_golang_v177_mongo/lib/env"
)

var Client *mongo.Client
var Ctx context.Context

func Mongodb() *mongo.Client {
	var err error
	var cancel context.CancelFunc
	Ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if Client, err = mongo.Connect(Ctx, options.Client().ApplyURI(env.String("DATABASE_DSN", "mongodb://localhost:27017"))); err != nil {
		log.Fatal(err)
	}
	return Client
}
