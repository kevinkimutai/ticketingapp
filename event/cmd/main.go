package main

import (
	"os"
	"strconv"

	"github.com/charmbracelet/log"

	"github.com/kevinkimutai/ticketingapp/event/adapters/db"
	"github.com/kevinkimutai/ticketingapp/event/adapters/grpc"
	"github.com/kevinkimutai/ticketingapp/event/application/api"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//GET ENV VARIABLES
	PORT := os.Getenv("PORT")
	DBURL := os.Getenv("DB_URL")

	//Convert Port to int
	portInt, err := strconv.Atoi(PORT)
	if err != nil {
		log.Fatal("Error converting port err")
	}

	dbAdapter, err := db.NewAdapter(DBURL)

	if err != nil {
		log.Fatal("couldnt connect to DB", err)
	}
	application := api.NewApplication(dbAdapter)

	grpcServer := grpc.NewAdapter(application, portInt)

	grpcServer.Run()
}
