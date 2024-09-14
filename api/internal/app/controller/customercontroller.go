package controller

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/korn1k/stripepayment/internal/app/domain/stripecustomer"
	"github.com/korn1k/stripepayment/internal/app/util"
)

type ApiCustomerController struct {}

func NewApiCustomerController() *ApiCustomerController {
    return &ApiCustomerController{}
}

func (acc ApiCustomerController) Handle(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Failed to parse form", http.StatusBadRequest)
        return
    }

    email := r.FormValue("email")
    if !util.IsValidEmail(email) {
        http.Error(w, "Failed to validate email", http.StatusBadRequest)
        return
    }

    name := r.FormValue("name")
    if name != "" {
        http.Error(w, "Failed to validate name", http.StatusBadRequest)
        return
    }

    cust, err := stripecustomer.NewCustomer(name, email)
    if err != nil {
        http.Error(w, "Failed to create customer", http.StatusInternalServerError)
        return 
    }

    // Note, we are doing this because the backend does not have a persisting storage.
    // !!! DO NOT DO THIS IN WORKING EXAMPLES !!!
    custRef := uuid.New()
    cookie := &http.Cookie{
        Name:     custRef.String(),
        Value:    cust.ID,
        Path:     "/",
        MaxAge:   3600, // 1 hour
        HttpOnly: true,
    }

    http.SetCookie(w, cookie)

    w.Write([]byte(custRef.String()))
}
