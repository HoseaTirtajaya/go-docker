package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetClient() *mongo.Client {
	var connectMongo string = os.Getenv("MONGO_CLIENT")
	clientOptions := options.Client().ApplyURI(connectMongo)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func main() {
	godotenv.Load(".env")
	c := GetClient()
	var isConnect bool = false

	err := c.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		isConnect = true
		log.Println("Connected!")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var name = os.Getenv("NAME")
		log.Println(os.Getenv("NAME"), os.Getenv("SECRET"))

		for isConnect == true {
			log.Println("Connected to mongodb")
			fmt.Fprintf(w, "Hello %s ", name)
		}
	})

	http.ListenAndServe(":8080", nil)
}
