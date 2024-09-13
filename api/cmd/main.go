package main

import (
	"fmt"
	"net/http"

	"github.com/korn1k/stripepayment/internal/config"
	"github.com/korn1k/stripepayment/internal/router"
	"github.com/stripe/stripe-go/v75"
)

func main() {
	router := router.NewRouter()

	c := config.MustInitConfig()
	stripe.Key = c.Stripe.SecretKey

	port := "8080"
	fmt.Printf("Starting server on port %s...\n", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Errorf("Error starting server: %v\n", err)
	}
}