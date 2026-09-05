package main

import "fmt"

func main() {
	/*
		## What is enums?
		- An enum (short for enumeration) is a special data type in computer programming
		  used to define a fixed, restricted set of named constants

		- Instead, idiomatic Go implements enums by defining a custom type
		  Go does not have a built-in enum type as a distinct language feature.
		  based on an integer or string and combining it with a const block
		  and the iota keyword.

		### What is iota?
		In Go, iota is a predefined or special identifier used within const blocks,
		that auto-increments starting at 0 within a const block.

		- we use enums to define type of order statuses, user roles and categories and many more.

	*/

	// we can define our custom type by below syntax
	authenticateRoleAccess(2)
	authenticateRoleAccess(SuperAdmin)
}

type userRole int // now userRole is indirectly a int type like you can define int type by using userRole

// we can group constants using ()
const (
	SuperAdmin userRole = iota // 0
	Admin
	Editor
	Commenter
	Viewer
)

func authenticateRoleAccess(role userRole) {
	if role == SuperAdmin {
		fmt.Println("You are a super admin, and have whole acess to the system.")
	} else if role == Admin {
		fmt.Println("Admin: You have the system access.")
	} else if role == Editor {
		fmt.Println("You have only edit access.")
	} else if role == Commenter {
		fmt.Println("You can only comment on the document")
	} else if role == Viewer {
		fmt.Println("You can only view the document.")
	}
}

// we can also define string type enums

type orderStatus string

const (
	Recieved   orderStatus = "Recieved"
	Processing orderStatus = "Processing"
	Pending    orderStatus = "Pending"
	Delivered  orderStatus = "Delivered"
)
