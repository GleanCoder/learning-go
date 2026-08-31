package main

import "fmt"

// by doing in this struct way we have call pay method on evey instance and we have to modify the code for every payment gateway.
// to solve this we can use interface.

type paymentGateway interface {
	pay(amount float32)
}

// type payment struct {
// }

// func (p payment) makePayment(amount float32) {
// 	razorpayPaymentGw := razaorpay{}
// 	razorpayPaymentGw.pay(amount)
// }

func checkOutPayment(gateway paymentGateway, amount float32) {
	gateway.pay(amount)

}

type razaorpay struct{}

func (r razaorpay) pay(amount float32) {
	// razorpay api call and logics
	fmt.Println("Payment of", amount, "done using Razorpay!")
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("Payment of", amount, "done using Stripe!")
}

func main() {
	/*
		An interface in Go defines a set of method signatures — no implementation, no fields.
		Any type that implements all those methods automatically satisfies the interface.
		 No implements keyword needed like other languages.
	*/

	//USE CASE: suppose we want to implement multiple payment gateway service, let's start it with struct then will move to interface.
	// newPayment := payment{}
	// newPayment.makePayment(199)

	razaorpayGW := razaorpay{}
	checkOutPayment(razaorpayGW, 399)

	stripeGw := stripe{}
	checkOutPayment(stripeGw, 149)

	/*
		here our payment checkout doesn't depends on any gateway and also any gateway which satisfy the method of our payment interface will automatically get access to our pay(), don't need to use implements keyword like other languages, go will do it automatically if methods signature is same.
		- by this we follow open for extension and close for modification rule.
	*/

}
