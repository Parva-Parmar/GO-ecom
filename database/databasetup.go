package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.Client {
    // Set client options
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

    // Connect to MongoDB
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    // Check the connection
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal("Failed to connect to MongoDB:", err)
        return nil
    }

    fmt.Println("Successfully connected to MongoDB")

    return client
}

// func DBSet() *mongo.Client {

// 		client,err := mongo.NewClient(options.Client().ApplyURI("mongodp://localhost:27017"))
// 		if err != nil{
// 			log.Fatal(err)
// 		}

// 		ctx , cancel := context.WithTimeout(context.Background() , 10*time.Second)
// 		defer cancel()

// 		err = client.Connect(ctx)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		client.Ping(context.TODO(),nil)
// 		if err != nil{
// 			log.Fatal("Failed to connect to mongodb")
// 			return nil
// 		}

// 		fmt.Println("Succesfully connect to mongodb")

// 		return client
// }

var Client *mongo.Client = DBSet()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
		var collection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)
		return collection
}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
		var productCollection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)
		return productCollection
}
 