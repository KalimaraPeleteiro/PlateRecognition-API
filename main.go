package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client   *mongo.Client
	Database *mongo.Database
}

var mg MongoInstance

func connect(databaseName, mongoURI string) error {
	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	db := client.Database(databaseName)

	mg = MongoInstance{
		Client:   client,
		Database: db,
	}

	log.Println("Conexão ao banco finalizada.")
	return nil
}

func main() {
	godotenv.Load()

	databaseName := os.Getenv("DATABASE")
	if databaseName == "" {
		log.Fatal("Não foi encontrada a variável de ambiente de banco.")
	}

	uri := os.Getenv("URI")

	if uri == "" {
		log.Fatal("Não foi encontrada a variável de ambiente de URI.")
	}

	if err := connect(databaseName, uri+databaseName); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	log.Fatal(app.Listen(":3000"))
}
