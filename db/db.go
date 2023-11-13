package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var dbname string

func setClient(c *mongo.Client) {
	client = c
}
func getClient() *mongo.Client {
	if client == nil {
		panic("Client not initialized")
	}
	return client
}
func setDBName(db string) {
	dbname = db
}

func Connect(uri string, dbname string) *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI))
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	log.Println("Connected to MongoDB!")
	setClient(client)
	setDBName(dbname)
	return client
}

func GetDB() *mongo.Database {
	if dbname == "" {
		panic("DB name not initialized")
	}
	return getClient().Database(dbname)
}

func Disconnect() {
	err := getClient().Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
}
