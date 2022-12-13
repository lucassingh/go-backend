package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN object to connect to database*/
var MongoCN = ConnectDB()

var clientOptions = options.Client().ApplyURI("mongodb+srv://goBack:QOZ1rT7iQ16TpUwU@go-backend.vu6qxpt.mongodb.net/test")

/*ConnectDB is the Conection to database */
func ConnectDB() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conected to database !!!!")

	return client
}

/*CheckConnection verify if the connection is ok or no*/
func CheckConnection() int {

	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {

		return 0
	}

	return 1
}
