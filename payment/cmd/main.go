package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/kevinkimutai/ticketingapp/payment/adapters/auth"
	"github.com/kevinkimutai/ticketingapp/payment/adapters/db"
	"github.com/kevinkimutai/ticketingapp/payment/adapters/grpc"
	stripepkg "github.com/kevinkimutai/ticketingapp/payment/adapters/stripe"
	"github.com/kevinkimutai/ticketingapp/payment/application/api"
	// "log"
	// "os"
	// "strconv"
	// "github.com/joho/godotenv"
	// "github.com/kevinkimutai/go-grpc/order/adapters/db"
	// "github.com/kevinkimutai/go-grpc/order/adapters/grpc"
	// "github.com/kevinkimutai/ticketingapp/auth/adapters/gateway"
	// "github.com/kevinkimutai/ticketingapp/auth/application/api"
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
	//HTTPPORT := os.Getenv("HTTP_PORT")
	STRIPEKEY := os.Getenv("STRIPE_KEY")

	//Convert Port to int
	portInt, err := strconv.Atoi(PORT)
	if err != nil {
		log.Fatal("Error converting port err")
	}

	//Convert HTTP_Port to int
	// httpPort, err := strconv.Atoi(HTTPPORT)
	// if err != nil {
	// 	log.Fatal("Error converting port err")
	// }

	dbAdapter, err := db.NewAdapter(DBURL)
	if err != nil {
		log.Fatal("couldnt connect to DB", err)
	}

	authAdapter, err := auth.NewAdapter(AUTHURL)
	if err != nil {
		log.Fatal("couldnt connect to Auth Service", err)
	}

	stripeAdapter := stripepkg.NewAdapter(STRIPEKEY)

	application := api.NewApplication(dbAdapter, stripeAdapter, authAdapter)

	grpcServer := grpc.NewAdapter(application, portInt)
	//gatewayServer := gateway.NewAdapter(portInt, httpPort)

	//go gatewayServer.Run()
	grpcServer.Run()
}

// package main

// import (
// 	"github.com/stripe/stripe-go/v76"
// 	"github.com/stripe/stripe-go/v76/checkout/session"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	// This is a public sample test API key.
// 	// Don’t submit any personally identifiable information in requests made with this key.
// 	// Sign in to see your own test API key embedded in code samples.
// 	stripe.Key = ""

// 	http.Handle("/", http.FileServer(http.Dir("public")))
// 	http.HandleFunc("/create-checkout-session", createCheckoutSession)
// 	addr := "localhost:4242"
// 	log.Printf("Listening on %s", addr)
// 	log.Fatal(http.ListenAndServe(addr, nil))
// }

// func createCheckoutSession(w http.ResponseWriter, r *http.Request) {
// 	domain := "http://localhost:4242"
// 	params := &stripe.CheckoutSessionParams{
// 		LineItems: []*stripe.CheckoutSessionLineItemParams{
// 			&stripe.CheckoutSessionLineItemParams{
// 				// Provide the exact Price ID (for example, pr_1234) of the product you want to sell
// 				Price:    stripe.String("{{PRICE_ID}}"),
// 				Quantity: stripe.Int64(1),
// 			},
// 		},
// 		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
// 		SuccessURL: stripe.String(domain + "/success.html"),
// 		CancelURL:  stripe.String(domain + "/cancel.html"),
// 	}

// 	s, err := session.New(params)

// 	if err != nil {
// 		log.Printf("session.New: %v", err)
// 	}

// 	http.Redirect(w, r, s.URL, http.StatusSeeOther)
// }
