package main

import "fmt"

func main() {

	/*
		- Struct embedding is used to reuse the fields and methods of one struct inside another struct without explicitly nesting it.
		- as we know there is no concept of inheritance in golang but suppose we want to use one struct and its methods in another different struct then we can use embedded struct.

	*/

	type customer struct {
		name string
		mob  int
	}

	type order struct {
		item     string
		quantity int
		price    float32
		customer
	}

	// we can assign customer inside order in 2 ways lets use  them

	//1st way: create customer separate and assign it to order

	customerOne := customer{
		name: "Aditya",
		mob:  7848000100,
	}

	pizzaOrder := order{
		item:     "Pizza",
		quantity: 1,
		price:    140,
		customer: customerOne,
	}

	fmt.Println(pizzaOrder)

	//2nd way create customer inside order itself

	pastaOrder := order{
		item:     "Pasta",
		quantity: 2,
		price:    600,
		customer: customer{name: "Kiran", mob: 7848000100},
	}

	fmt.Println(pastaOrder)

}
