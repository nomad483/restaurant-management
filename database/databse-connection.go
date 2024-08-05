package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var Client *mongo.Client = Connect()

func Connect() *mongo.Client {
	fmt.Println("trying to connect to mongodb")

	mongoURI := os.Getenv("MONGODB_URI")

	fmt.Println(mongoURI)

	if mongoURI == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	//Client = client

	//defer func() {
	//	if err = client.Disconnect(context.TODO()); err != nil {
	//		panic(err)
	//	}
	//}()

	fmt.Println("Connected to MongoDB!")
	return client

}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)

	return collection
}
