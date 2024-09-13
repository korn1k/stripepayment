package stripecustomer

import (
	"github.com/stripe/stripe-go/v75/customer"
	"github.com/stripe/stripe-go/v75"
)

func NewCustomer(name, email string) (*stripe.Customer, error) {
	params := &stripe.CustomerParams{
		Email: stripe.String(email),
		Name: stripe.String(name),
	}

	cust, err := customer.New(params)
	if err != nil {
		return nil, err
	}

	return cust, err
}