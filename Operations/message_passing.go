package main

import "fmt"

type CreditAccount struct{}

func (nesne *CreditAccount) processPayment(amount float32) {
	fmt.Println("islem...")
}

func Credit(value chan float32) *CreditAccount {

	temp := &CreditAccount{}

	go func(ch chan float32) {
		for amount := range ch {
			CreditAccount.processPayment(amount)
		}
	}(ch)

	return CreditAccount
}

func main() {
	mych := make(chan float32)
	Credit(mych)
	mych <- 500
	var a string
	fmt.Sscanln(&a)
}

// Fatal example
