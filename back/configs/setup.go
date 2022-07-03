package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	usernameDb, passDb := EnvDB()

	uriDb := fmt.Sprintf("mongodb+srv://%s:%s@finalproject.hmjauoz.mongodb.net/?retryWrites=true&w=majority", usernameDb, passDb)

	// Set client options
	clientOptions := options.Client().ApplyURI(uriDb)

	// Cancel the connection if exceeds 10s.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	return client
}

//Client instance
var DB *mongo.Client = ConnectDB()

//getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("finalProject").Collection(collectionName)
	return collection
}
