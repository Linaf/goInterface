package payments

import (
	"errors"
	"fmt"
)

type PaymentOption interface {
	ProcessPayment(float32) bool
}

type CreditCard struct {
	ownerName       string
	cardNumber      string
	availableCredit float32
}

//constructor
func CreateCreditCardAcount(ownerName, cardNumber string) *CreditCard {
	return &CreditCard{
		ownerName:  ownerName,
		cardNumber: cardNumber,
	}

}

func (c CreditCard) OwnerName() string {
	return c.ownerName
}

func (c *CreditCard) SetOwnerName(value string) error {
	if len(value) == 0 {
		return errors.New("Invalid ownerName provided.")
	}
	c.ownerName = value
	return nil

}

func (c CreditCard) CardNumber() string {
	return c.cardNumber
}

func (c *CreditCard) SetCardNumber(value string) error {
	if len(value) == 0 {
		return errors.New("Invalid card number provided.")
	}
	c.cardNumber = value
	return nil

}

func (c CreditCard) AvailableCredit() float32 {
	return 5000.00
}

func (c CreditCard) ProcessPayment(float32) bool {
	fmt.Println("Processing Credit card payment ......")
	return true
}
