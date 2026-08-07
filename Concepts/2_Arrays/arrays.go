package main

import "fmt"

// Arrays are fixed-size collections of elements of the same type.
func main() {

	/*
		below is the syntax to declare an array

		- var or const variable_name [no. of elements] variable_type
		- to initialize value use array_variable_name[index_no.]= value

		- here our array size is 5 and our array type is int, but we initalized only one valuue at index 0 so the rest will be 0 and only 0 indexed array value will be 7
		- in case of string it will be empty string
		- in case of boolean it will be false
		- in case of float it will be also 0

	*/

	var numArray [5]int

	numArray[0] = 7

	fmt.Println(numArray[0])
	fmt.Println(numArray)

	// we can get the length of array using len() built-in function, => len(array_name)

	fmt.Println(len(numArray))

	// we can also declare and initialize in a single line

	// using short hand notation variableName := [size]type{values}

	fullName := [3]string{"Aditya", "Kiran", "Acharya"}
	fmt.Println(fullName)

	var stateCode = [2]string{"OD", "AP"}
	fmt.Println(stateCode)

	// we can also declare 2D array by below way:

	setValue := [2][2]int{{1, 2}, {2, 3}}

	fmt.Println(setValue)

	/*
		- there also another way which similar to array in go, and that is slice.
		- it's recomended use array when you know the length of the array, otherwise use slice and also slice has a topic called range which we will do in slice folder
		- basically Use arrays when the size is fixed and known at compile time. Otherwise, use slices.

	*/

	/*
		in array we can assign an array of value to another, by this go create a separate copy of the array and assign it to the other
		- by this change in other will not affect the another
	*/

	myName := fullName

	myName[1] = "Shekhar"

	fmt.Println(fullName) // output: [Aditya Kiran Acharya]
	fmt.Println(myName)   // output: [Aditya Shekhar Acharya]

	/*
	 NOTE: But in slice both refer to the same underlying array, so changing one will be visible through the other.

	*/
}
