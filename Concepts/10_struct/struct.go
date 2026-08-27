package main

import (
	"fmt"
	"time"
)

// suppose If we want the method to modify the original struct:

type bankAccount struct {
	name     string
	bankName string
	balance  float32
	active   bool
}

// we have to add a pointer reciever, which will help us to modify the struct
// convention rule of reciever: first a single character from the struct name's first char then struct name, how we written below

func (b *bankAccount) debitAmount(amount float32) float32 {
	b.balance = b.balance - amount
	return b.balance
}

// if you dont want to modify anything then use simple receiever type instead of pointer receiever type.

// CONSTRUCTOR:

/*
We use a constructor with a struct to create and initialize it consistently, optionally applying default values or validation.

In Go, it's usually a NewType() function, not a special constructor keyword.

- it's convention rule to use New word to create constructor function in golang newconstructorVarName


*/

func NewBankAccount(name string, bankName string, balance float32) *bankAccount {
	myBankAccount := bankAccount{
		name:     name,
		bankName: bankName,
		balance:  balance,
		active:   true,
	}

	return &myBankAccount
}

func main() {
	/*
		- In Go (Golang), a struct (short for structure) is a user-defined,
		  composite data type that allows you to group related fields of different data types into a single unit.

		- Because Go does not feature traditional classes or inheritance,
		  structs serve as the primary tool for creating blueprints, implementing object-oriented patterns,
		  and modeling real-world entities.

	*/

	type order struct {
		id        string
		item      string
		price     float32
		status    string
		createdAt time.Time // we have a time package and it has nanosecond precesion
	}

	customerOrder1 := order{
		id:     "100374s34ya",
		item:   "Biriyani",
		price:  160.00,
		status: "recieved",
	}

	// suppose we want to add  or modify a value later then we can use .

	customerOrder1.createdAt = time.Now()
	fmt.Println(customerOrder1)

	// to get any particular object from struct we can use the same .

	fmt.Println(customerOrder1.price)

	// we can use the struct name to create multiple instance of the struct, here above how we have used order to create an instance and each instance will be different from each other.
	// if you don't set any field value, then default value will be zero value.

	customerOne := bankAccount{
		name:     "Tony",
		bankName: "Stark Finance",
		balance:  20000.00,
		active:   true,
	}
	fmt.Println(customerOne)

	// this is how we call the method below

	customerOne.debitAmount(3200)

	fmt.Println("balance after money debited:", customerOne.balance)

	customerTwo := NewBankAccount("Aditya", "SBI", 200000)
	fmt.Println(customerTwo)

	// we are creating blueprint using struct then  we are making lots of instance using that struct.
	// what if we want to use a struct only once ? for that we can follow the below syntax, it's quite similar to how we create object in typescript using object literals

	language := struct {
		name   string
		isGood bool
	}{"Golang", true}
	fmt.Println(language)
}
