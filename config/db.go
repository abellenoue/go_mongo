package config

import (
	"cesi/go_mongo/controllers"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connect() {
	// Database Config
	// clientOptions := options.Client().ApplyURI("mongodb://user:password@host:port/test?authSource=admin&replicaSet=Cluster0-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true")
	clientOptions := options.Client().ApplyURI("mongodb+srv://Admin:%40AZERTY12@mycluster.xbgky.mongodb.net/test")
	client, err := mongo.NewClient(clientOptions)

	//Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	//To close the connection at the end
	defer cancel()

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	db := client.Database("go_mongo")
	controllers.TodoCollection(db)
	return
}

func TestConnect() {
	// Database Config
	// clientOptions := options.Client().ApplyURI("mongodb://user:password@host:port/test?authSource=admin&replicaSet=Cluster0-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true")
	clientOptions := options.Client().ApplyURI("mongodb+srv://Admin:%40AZERTY12@mycluster.xbgky.mongodb.net/test")
	client, err := mongo.NewClient(clientOptions)

	//Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	//To close the connection at the end
	defer cancel()

	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	db := client.Database("go_mongo_test")
	controllers.TodoCollection(db)
	return
}
