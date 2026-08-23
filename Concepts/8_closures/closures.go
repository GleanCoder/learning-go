package main

import "fmt"

/*
- closure in golang is the function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
- if inner function is dependent on the outer function's variables, then at runtime the binding  of the outer function's variables to the inner function is maintained, even after the outer function has returned. This is because the inner function maintains a reference to the outer function's variables, allowing it to access and modify them even after the outer function has completed execution.

*/

func main() {
	usedAmount := allocatedBudget()
	fmt.Println(usedAmount(200)) //800
}

func allocatedBudget() func(amount int) int {
	budgetAmount := 1000

	return func(amount int) int {
		budgetAmount = budgetAmount - amount
		return budgetAmount
	}
}
