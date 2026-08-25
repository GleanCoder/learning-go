package main

import "fmt"

/*
POINTERS:

- pointers is the variable that stores memory address of another variable.
- pointers are used to indirectly refer to the value stored in a variable, rather than the value itself.
	- basically this line is distinguishing two ways of working with a variable's data: directly vs. indirectly.

	1. Direct access: you use the variable itself and get its value straight away.
	2. Indirect access (via a pointer): instead of holding the value, a pointer holds the address where that value lives in memory.
	   To actually see or change the value, you have to go through the pointer — you're referring to it "indirectly,".

- to declare a pointer we have to use * along with the data_type of the variable it going to point
- and to initialize pointer we have to use &variable_name (&- this is called address of operator)





*/

func main() {

	/*
		- var variable_name *data_type : this is how we declare a pointer in golang.
			- here the data type will be assign to the value which is stored in the data block,  which data block's address we have stored via pointer.
	*/

	num := 2

	var ptr *int = &num

	fmt.Println(ptr)  // address of num
	fmt.Println(*ptr) // value inside num data block

	// if we declare a pointer variable and don't assign any value to it then its default value will be nil.

	// short way of intialize pointer

	secondPointer := &num
	// by the above syntax we can store any type of value

	fmt.Println(secondPointer)

	changeValueByRef(ptr)

	fmt.Println(num) // value == 8

}

/*
Because Go is fundamentally a pass-by-value language, every time you pass a variable into a function or method,
Go creates a brand-new copy of that data in memory.
Pointers allow you to pass the memory address of the original data instead of a copy.
*/

func changeValueByRef(num *int) {
	*num = *num + 6
}
