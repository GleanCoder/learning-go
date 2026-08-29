package main

import "fmt"

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Payment of", amount, "is done via Razorpay!")
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("Payment of", amount, "is done via Stripe!")
}

func main() {
	/*
		An interface in Go defines a set of method signatures — no implementation, no fields.
		Any type that implements all those methods automatically satisfies the interface.
		 No implements keyword needed like other languages.
	*/

	//USE CASE: suppose we want to implement multiple payment gateway service, let's start it with struct then will move to interface.

	razorpayGw := razorpay{}
	razorpayGw.pay(100)

	stripeGw := stripe{}
	stripeGw.pay(200)
}
