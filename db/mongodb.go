package db

import (
	"go-katrade/configs"
	"log"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() *mongo.Client{
	client, err := mongo.NewClient(options.Client().ApplyURI(configs.GetEnv("MONGOURI")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil{
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to MongoDB")
    return client
}

var DB *mongo.Client = ConnectMongoDB()