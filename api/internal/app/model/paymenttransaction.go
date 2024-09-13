package model

import "github.com/google/uuid"

type PaymentMethodType string

const (
	Card string = "card"
)

type PaymentTransaction struct {
	ID uuid.UUID
	Funds Money
	PaymentMethodTypes []string
	StripeCustomerID string
	StripePaymentIntentID string
}

func NewPaymentTransaction(id uuid.UUID, funds Money, paymentMethodTypes []string) *PaymentTransaction {
	return &PaymentTransaction{ID: id, Funds: funds, PaymentMethodTypes: paymentMethodTypes}
}

func (pt PaymentTransaction) HasStripeCustomerID() bool {
	return pt.StripeCustomerID != ""
}

func (pt PaymentTransaction) HasPaymentMethodTypes() bool {
	return len(pt.PaymentMethodTypes) > 0
}

func (pt PaymentTransaction) HasStripePaymentIntentID() bool {
	return pt.StripePaymentIntentID != ""
}

type Money struct {
	Amount int64
	Currency string
}