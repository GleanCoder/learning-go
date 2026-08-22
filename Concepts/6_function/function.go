package main

import "fmt"

func main() {

	/*

			Go function syntax:

			func name(parameters) returnType {
		    ...
		}

		- in Go Main func is special func or the starting point of execution, so we don't need to call it explicitly


	*/

	fmt.Println(sumOfNum(5, 10))

	fmt.Println(getLanguages())

	// we can also get the return values in below way:

	language1, language2, _ := getLanguages()

	fmt.Println(language1, language2)

	// suppose one value we don't want to use then we can use _ while getting the value to suppress the error.

	fn := func(a int) int {
		return 2
	} // here we have defined an anonymous function.

	fmt.Println(processIt(fn))

	fnInt := returnFunction()
	fmt.Println(fnInt(6))

}

/*
in function we have to define return type, if we don't specify the return type then it will throw us the below error

too many return values
	have (int)
	want ()

- want () — based on your function signature, Go expects zero return values (empty parens = nothing)
- have (int) — but your return sum statement is trying to send back one int value
-too many return values — you're returning more values than the signature promised

*/

func sumOfNum(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

/*
- suppose we want to return multiple values in our function, in that case we have to group the return types under ()- parenthesis.
- and also the order of return type which defined inside () must be follow while return values.
*/

func getLanguages() (string, string, bool) {
	return "golang", "javascript", true
}

/*
- In golang functions consider as first class citizen, that means:
	- we can assign functions to a variable.
	- we can pass one function as an argument inside another function.
	- we can also return a new function from another function.

*/

// function taking another function as an argument
func processIt(fn func(a int) int) int {
	return fn(1)
}

// function returning another function
func returnFunction() func(a int) int {
	return func(a int) int {
		return a
	}
}
