package controller

import "net/http"

type ApiPaymentIntentController struct {}

func NewApiPaymentIntentController() *ApiPaymentIntentController {
    return &ApiPaymentIntentController{}
}

func (apic ApiPaymentIntentController) Handle(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
        http.Error(w, "Failed to parse form", http.StatusBadRequest)
        return
    }

	custRef := r.FormValue("customerReference")
	
	cookie, err := r.Cookie(custRef)
    if err != nil {
        if err == http.ErrNoCookie {
            http.Error(w, "Cookie not found", http.StatusNotFound)
        } else {
            http.Error(w, "Error retrieving cookie", http.StatusInternalServerError)
        }

        return
    }

    amount := r.FormValue("amount")
	currency := r.FormValue("currency")
}