package main

import (
	"os"
	"strconv"

	"github.com/charmbracelet/log"

	"github.com/kevinkimutai/ticketingapp/event/adapters/auth"
	"github.com/kevinkimutai/ticketingapp/event/adapters/db"
	"github.com/kevinkimutai/ticketingapp/event/adapters/gateway"
	"github.com/kevinkimutai/ticketingapp/event/adapters/grpc"
	"github.com/kevinkimutai/ticketingapp/event/adapters/organiser"
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
	AUTHURL := os.Getenv("AUTH_URL")
	ORGANISERURL := os.Getenv("ORGANISER_URL")
	HTTPPORT := os.Getenv("HTTP_PORT")

	//Convert Port to int
	portInt, err := strconv.Atoi(PORT)
	if err != nil {
		log.Fatal("Error converting port err")
	}

	//Convert HTTP_Port to int
	httpPort, err := strconv.Atoi(HTTPPORT)
	if err != nil {
		log.Fatal("Error converting port err")
	}

	dbAdapter, err := db.NewAdapter(DBURL)
	if err != nil {
		log.Fatal("couldnt connect to DB", err)
	}

	authAdapter, err := auth.NewAdapter(AUTHURL)
	if err != nil {
		log.Fatal("couldnt connect to Auth Service", err)
	}

	organiserAdapter, err := organiser.NewAdapter(ORGANISERURL)
	if err != nil {
		log.Fatal("couldnt connect to Organiser Service", err)
	}

	application := api.NewApplication(dbAdapter, organiserAdapter)

	grpcServer := grpc.NewAdapter(application, portInt, authAdapter)
	gatewayServer := gateway.NewAdapter(portInt, httpPort)

	go gatewayServer.Run()
	grpcServer.Run()
}
