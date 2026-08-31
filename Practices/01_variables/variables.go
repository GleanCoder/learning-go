package main

import "fmt"

func main() {
	/* 1. Declare an int, float64, string, and bool using var,
	then declare the same four using :=.
	Print all of them with %T to confirm types.
	*/

	// var num int = 1
	// var amount float64 = 100.54
	// var name string = "Kurkure"
	// var isLoggedIn = true

	num := 1
	amount := 99.99
	name := "Lays"
	isLoggedIn := false

	fmt.Printf("type of num is %T\n", num)
	fmt.Printf("type of amount is %T\n", amount)
	fmt.Printf("type of name is %T\n", name)
	fmt.Printf("type of isLoggedIn is %T\n", isLoggedIn)

	/*
		2. Write code that swaps two variables a, b := 5, 10 without using a
		temp variable (Go's multi-assignment trick).
	*/
	a, b := 5, 10
	a, b = b, a
	fmt.Println(a, b)

	//Q3.

	for d := Sunday; d <= Saturday; d++ {
		fmt.Println(d)
	}

	//4. Show what happens when you print an uninitialized var x int, var s string, var b bool — confirm zero values without setting them.

	var x int
	var s string
	var boolValue bool
	fmt.Println(x, s, boolValue) // 0  false

}

/*
3. Declare a constant block with iota for days of the week (Sunday=0 ... Saturday=6), print each.
*/
type weekDays int

const (
	Sunday weekDays = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (d weekDays) days() string {
	return [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}[d]
}
