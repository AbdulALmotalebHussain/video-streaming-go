package database

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
)

var Client *mongo.Client

func InitMongoDB(uri string) {
    clientOptions := options.Client().ApplyURI(uri)
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }
    if err = client.Ping(context.Background(), nil); err != nil {
        log.Fatal(err)
    }
    Client = client
    log.Println("Connected to MongoDB!")
}

