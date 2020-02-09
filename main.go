package main

import "github.com/Linaf/goInterface/payments"

func main() {

	var option payments.PaymentOption
	option = payments.CreateCreditCardAcount("lina", "cn12445")
	option.ProcessPayment(200)
	option = payments.CreateCashAccount()
	option.ProcessPayment(300)

}
