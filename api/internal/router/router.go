package router

import (
	"net/http"

	"github.com/korn1k/stripepayment/internal/app/controller"
)

func NewRouter() *http.ServeMux {
    router := http.NewServeMux()

	apiCustomerController := controller.NewApiCustomerController()
    router.HandleFunc("/customer/new", apiCustomerController.Handle)

	healthController := controller.NewHealthController()
    router.HandleFunc("/health", healthController.Handle)

    return router
}