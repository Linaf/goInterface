package payments

import "fmt"

type Cash struct{}

func CreateCashAccount() *Cash {
	return &Cash{}
}

func (c Cash) ProcessPayment(float32) bool {
	fmt.Println("Processing cash payment.")
	return true
}
