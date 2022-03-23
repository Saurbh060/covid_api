package database

import (
	m "covid_api/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(mongoFiled []interface{}) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")
	defer client.Disconnect(ctx)

	database := client.Database("covidCases")
	stateCollection := database.Collection("stateData")

	fmt.Println("printing mongo data")
	fmt.Println(mongoFiled)

	insertManyResult, err := stateCollection.InsertMany(ctx, mongoFiled)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
	fmt.Println("Connection to MongoDB closed.")
}

func ConnectAndGet(stateName string) m.MongoFields {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")
	defer client.Disconnect(ctx)

	database := client.Database("covidCases")
	stateCollection := database.Collection("stateData")

	var result m.MongoFields
	if err = stateCollection.FindOne(ctx, bson.M{"state": stateName}).Decode(&result); err != nil {
		log.Fatal(err)
	}

	fmt.Println("printing podcast ", result)

	// for i := 0; i < len(podcast.StateCases); i++ {
	// 	// fmt.Println("key", podcast.StateCases[i])
	// 	data := podcast.StateCases[i]
	// 	// fmt.Println("state:", data.State)
	// 	if data.State == stateName {
	// 		result.State = data.State
	// 		result.TotalCases = data.TotalCases
	// 		break
	// 	}
	// }
	fmt.Println("Connection to MongoDB closed.")
	return result
}
