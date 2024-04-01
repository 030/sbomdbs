package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Define a struct to represent your JSON data
type Person struct {
	Name  string `bson:"name"`
	Age   int    `bson:"age"`
	Email string `bson:"email"`
}

func main() {
	// Set up MongoDB client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:9999")

	// Set username and password as options
	clientOptions.Auth = &options.Credential{
		Username: "mongoadmin",
		Password: "secret",
	}

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Ping the MongoDB server to verify the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// Get a handle for your "test" database
	database := client.Database("test")

	// Get a handle for your "people" collection
	collection := database.Collection("people")

	// // Create a sample JSON object
	// person := Person{Name: "John Doe", Age: 30, Email: "john@example.com"}

	// // Convert struct to JSON bytes
	// personBson, err := bson.Marshal(person)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Insert JSON data into MongoDB
	// insertResult, err := collection.InsertOne(context.Background(), personBson)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Inserted document with ID:", insertResult.InsertedID)

	// // Open the file
	// file, err := os.ReadFile("data.json")
	// if err != nil {
	// 	log.Fatal("Error reading file:", err)
	// }

	// // Define an interface to store the JSON data
	// var data []map[string]interface{}

	// // Unmarshal the JSON data into the interface
	// if err := json.Unmarshal(file, &data); err != nil {
	// 	log.Fatal("Error unmarshaling JSON data:", err)
	// }

	// // Insert each document into MongoDB
	// for _, doc := range data {
	// 	// insertResult, err := collection.InsertOne(context.Background(), doc)
	// 	_, err := collection.InsertOne(context.Background(), doc)
	// 	if err != nil {
	// 		log.Fatal("Error inserting document into MongoDB:", err)
	// 	}

	// 	// fmt.Println("Data inserted into MongoDB!")
	// 	// // Read the inserted document
	// 	// var result Person
	// 	// err = collection.FindOne(context.Background(), bson.M{"_id": insertResult.InsertedID}).Decode(&result)
	// 	// if err != nil {
	// 	// 	log.Fatal("Error reading document from MongoDB:", err)
	// 	// }

	// 	// fmt.Println("Read document from MongoDB:", result)
	// }

	// Read JSON data from file
	data, err := os.ReadFile("data.json")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	// Parse JSON data into BSON
	var bsonData bson.M
	err = bson.UnmarshalExtJSON(data, true, &bsonData)
	if err != nil {
		log.Fatal("Error unmarshaling JSON data:", err)
	}
	// bsonData["newDateField"] = time.Now()

	// Insert BSON document into MongoDB
	_, err = collection.InsertOne(context.Background(), bsonData)
	if err != nil {
		log.Fatal("Error inserting document into MongoDB:", err)
	}

	fmt.Println("Data inserted into MongoDB!")

	// // Define a filter to query the recently inserted document
	// filter := bson.M{"a": "b"}
	// // Find the document that matches the filter
	// var result bson.M
	// err = collection.FindOne(context.Background(), filter).Decode(&result)
	// if err != nil {
	// 	log.Fatal("Error finding document in MongoDB:", err)
	// }

	// fmt.Println("Inserted document:", result)

	// Find all documents in the collection
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal("Error finding documents in MongoDB:", err)
	}
	defer cursor.Close(context.Background())

	// Iterate over the cursor and print each document
	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal("Error decoding document:", err)
		}
		fmt.Println("Document:", result)
	}

	// Check for cursor errors after iteration
	if err := cursor.Err(); err != nil {
		log.Fatal("Cursor error:", err)
	}
}
