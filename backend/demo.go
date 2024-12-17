package main

import (
	"context"
	"fmt"
	"log"
	// "bson"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	uri := "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	
	client.Connect(context.Background())
	database := client.Database("demo")
	collection := database.Collection("demo")
	document := bson.D{
		{Key: "name", Value: "Asik"},
		{Key: "Roll", Value: "223CS1107"},
	}
	insertResult, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted document with ID: %v\n", insertResult.InsertedID)
	fmt.Println("Connected to MongoDB!")
	filter := bson.D{{Key: "name", Value: "Asik"}}

	
	var result bson.D 

	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal("Error finding document:", err)
	}

	fmt.Println("Found document:", result)
}