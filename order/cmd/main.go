package main

import (
	"os"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"github.com/kevinkimutai/ticketingapp/order/adapters/auth"
	"github.com/kevinkimutai/ticketingapp/order/adapters/db"
	"github.com/kevinkimutai/ticketingapp/order/adapters/gateway"
	"github.com/kevinkimutai/ticketingapp/order/adapters/grpc"
	"github.com/kevinkimutai/ticketingapp/order/adapters/payment"
	"github.com/kevinkimutai/ticketingapp/order/application/api"
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
	PAYMENTURL := os.Getenv("PAYMENT_URL")
	HTTPPORT := os.Getenv("HTTP_PORT")
	AUTHURL := os.Getenv("AUTH_URL")

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

	paymentAdapter, err := payment.NewAdapter(PAYMENTURL)
	if err != nil {
		log.Fatal("couldnt connect to Auth Service", err)
	}

	authAdapter, err := auth.NewAdapter(AUTHURL)
	if err != nil {
		log.Fatal("couldnt connect to Auth Service", err)
	}

	application := api.NewApplication(dbAdapter, paymentAdapter, authAdapter)

	grpcServer := grpc.NewAdapter(application, portInt)
	gatewayServer := gateway.NewAdapter(portInt, httpPort)

	go gatewayServer.Run()
	grpcServer.Run()
}
