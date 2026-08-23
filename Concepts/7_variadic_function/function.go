package main

import "fmt"

func main() {

	fmt.Println(summation(3, 4, 5, 6))

	// we can also provide slice as an argument for numbers, see  below

	nums := []int{3, 4, 5, 6}

	fmt.Println("Sum of slice elements:", summation(nums...))
}

/*
- variadic function is the function which can take any number of arguments.

How it works internally: inside the function, it is just a regular []variable_type slice. Go collects whatever arguments you pass into that slice automatically.

Rules:

Only the last parameter in a function can be variadic — you can't have two variadic params, and nothing can come after it
You can mix regular params with one variadic one at the end

*/

func summation(num ...int) int {
	sum := 0
	for _, value := range num {
		sum = sum + value
	}

	return sum
}
