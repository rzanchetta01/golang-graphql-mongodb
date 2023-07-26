package infra

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func MongoDbConnect() mongo.Client {

	//connection settings
	clientOptions := options.Client().ApplyURI(`mongodb+srv://admin:scratchdbapifakepassword@dbrpg.ttbqc.mongodb.net/?retryWrites=true&w=majority`)

	//connect
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("ERROR CONNECTING TO DB INSTANCE -> ", err)
	}

	//check connection
	conStatus, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(conStatus, nil)

	//Close connection if offline
	if err != nil {
		log.Fatal("CONNECTION OFFLINE -> ", err)
	}

	log.Println("CONNECTION OK!")
	return *client
}

func MongoDBDisconnect() {
	if client != nil {
		client.Disconnect(context.Background())
	}
}
