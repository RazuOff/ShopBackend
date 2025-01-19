package main

import (
	"log"

	"ginexample.com/pkg/api"
	postgres "ginexample.com/pkg/db/postgre"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load("../env/config.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	postgres.InitDB()
	api.StartServer(":8080")
}
