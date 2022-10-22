package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/harikesh-yadav/gofiber_api/database"
	"github.com/harikesh-yadav/gofiber_api/routes"
	"github.com/harikesh-yadav/gofiber_api/utils"
	"github.com/joho/godotenv"
)

func main() {

	loadEnv()

	app := fiber.New(fiber.Config{
		// Prefork: true,
	})

	url, err := utils.ConnectionUrlBuilder("fiber")
	if err != nil {
		log.Fatal(fmt.Errorf("not get server url %w", err))
		os.Exit(1)
	}

	_, err = database.Connection()

	if err != nil {
		log.Fatal("DB Not connected with error %w", err)
		os.Exit(1)
	} else {
		log.Println("DB connected")
	}

	routes.User(app)

	app.Listen(url)
}

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}

}
