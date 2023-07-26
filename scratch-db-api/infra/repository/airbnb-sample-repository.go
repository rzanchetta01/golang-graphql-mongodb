package repository

import (
	"context"
	"log"
	"scratch-db-api/entities"
	"scratch-db-api/infra"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionName = "listingsAndReviews"
var databaseName = "sample_airbnb"
var client mongo.Client = infra.MongoDbConnect()
var result []entities.Model

func GetAll() []entities.Model {
	collection := client.Database(databaseName).Collection(collectionName)
	filter := bson.D{}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal("Error retrieving sample document -> ", err)
	}

	for cursor.Next(context.Background()) {
		var i entities.Model
		err := cursor.Decode(&i)
		if err != nil {
			log.Println("Error decoding document -> ", err)
			continue
		}

		result = append(result, i)
	}

	log.Println("Get All Called")
	return result
}
