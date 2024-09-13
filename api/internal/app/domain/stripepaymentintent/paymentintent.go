package stripepaymentintent

import (
	"fmt"

	"github.com/korn1k/stripepayment/internal/app/model"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/paymentintent"
)

func NewPaymentIntent(transaction *model.PaymentTransaction) (*stripe.PaymentIntent, error) {
	if !transaction.HasPaymentMethodTypes() {
		return nil, fmt.Errorf("this payment intent requires payment method type specified")
	}

	if !transaction.HasStripeCustomerID() {
		return nil, fmt.Errorf("this payment intent requires customer ID specified")
	}

	if !transaction.HasStripePaymentIntentID() {
		return nil, fmt.Errorf("this payment intent requires payment intent ID specified")
	}

	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(transaction.Funds.Amount),
		Currency: stripe.String(transaction.Funds.Currency),
		Customer: stripe.String(transaction.StripeCustomerID),
		PaymentMethodTypes: stripe.StringSlice(transaction.PaymentMethodTypes),
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		return nil, err
	}

	return pi, nil
}