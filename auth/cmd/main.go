package main

import (
	"os"
	"strconv"

	"github.com/charmbracelet/log"

	"github.com/kevinkimutai/ticketingapp/auth/adapters/db"
	"github.com/kevinkimutai/ticketingapp/auth/adapters/gateway"
	"github.com/kevinkimutai/ticketingapp/auth/adapters/grpc"
	"github.com/kevinkimutai/ticketingapp/auth/application/api"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//GET ENV VARIABLES
	PORT := os.Getenv("GRPC_PORT")
	GATEWAYPORT := os.Getenv("HTTP_PORT")
	DBURL := os.Getenv("DB_URL")

	//Convert Port to int
	portInt, err := strconv.Atoi(PORT)
	if err != nil {
		log.Fatal("Error converting port err")
	}

	//Convert Gateway Port to int
	gatewayPortInt, err := strconv.Atoi(GATEWAYPORT)
	if err != nil {
		log.Fatal("Error converting port err")
	}

	dbAdapter, err := db.NewAdapter(DBURL)

	if err != nil {
		log.Fatal("couldnt connect to DB", err)
	}
	application := api.NewApplication(dbAdapter)

	grpcServer := grpc.NewAdapter(application, portInt)

	gatewayServer := gateway.NewAdapter(portInt, gatewayPortInt)

	go gatewayServer.Run()
	grpcServer.Run()
}
