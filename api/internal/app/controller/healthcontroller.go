package controller

import "net/http"

type HealthController struct {}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (hc HealthController) Handle(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}