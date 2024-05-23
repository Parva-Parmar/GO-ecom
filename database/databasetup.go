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

		client,err := mongo.NewClient(options.Client().ApplyURI("mongodp://localhost:27017"))
		if err != nil{
			log.Fatal(err)
		}

		ctx , cancel := context.WithTimeOut(context.Background() , 10*time.Second)
		defer cancel()

		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}

		client.Ping(context.TODO(),nil)
		if err != nil{
			log.Fatal("Failed to connect to mongodb")
			return nil
		}

		fmt.Println("Succesfully connect to mongodb")

		return client
}

		var Client *mongo.Client = DBSet()
		
func Userdata(client *mongo.Client, collectionName string) *mongo.Collection {

}

func ProductData(client *mongo.Client, collectionName string) *mongo.Clooection {

}
 