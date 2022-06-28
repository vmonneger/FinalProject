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

// // CreateIndex - creates an index for a specific field in a collection
// func CreateIndex(client *mongo.Client, collectionName string, field string, unique bool) bool {

// 	// 1. Lets define the keys for the index we want to create
// 	mod := mongo.IndexModel{
// 		Keys:    bson.M{"field": 1}, // index in ascending order or -1 for descending order
// 		Options: options.Index().SetUnique(unique),
// 	}

// 	// 2. Create the context for this operation
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	// 3. Connect to the database and access the collection
// 	collection := client.Database("finalProject").Collection(collectionName)

// 	// 4. Create a single index
// 	_, err := collection.Indexes().CreateOne(ctx, mod)
// 	if err != nil {
// 		// 5. Something went wrong, we log it and return false
// 		fmt.Println(err.Error())
// 		return false
// 	}

// 	// 6. All went well, we return true
// 	return true
// }

// var initIndex = CreateIndex(DB, "user", "email", true)
