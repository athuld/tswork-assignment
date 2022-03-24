package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getDBURI() string {
	err := godotenv.Load()
	checkError(err)

	return os.Getenv("DATABASE_URI")
}

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(getDBURI()))
	checkError(err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	checkError(err)

	err = client.Ping(ctx, nil)
	checkError(err)

	fmt.Print("Connected to Database")
	return client
}

//Client instance
var DB *mongo.Client = ConnectDB()

//getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("Tswork").Collection(collectionName)
	return collection
}
