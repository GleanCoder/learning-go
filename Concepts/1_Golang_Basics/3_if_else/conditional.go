package main

import "fmt"

func main() {
	age := 24

	if age >= 18 {
		fmt.Println("You're an adult now!")
	} else if age < 18 {
		fmt.Println("Oops! You're still a Minor.")
	}

	guessNumType(12)

	checkLargestNum(12, 18, 41)
}

/*
1. Even or Odd
Take a number and print whether it's even or odd.
*/

func guessNumType(num int) {
	if num%2 == 0 {
		fmt.Println(num, "is even no.")
	} else {
		fmt.Println(num, "is an odd no.")
	}
}

/*
2. Largest of Three
Take three numbers a, b, c and print the largest one.
*/

func checkLargestNum(num1 int, num2 int, num3 int) {
	if num1 > num2 && num1 > num3 {
		fmt.Println(num1, "is greater than other 2 numbers.")
	} else if num2 > num3 {
		fmt.Println(num2, "is greater than other 2 numbers.")
	} else {
		fmt.Println(num3, "is greater than other 2 numbers.")
	}
}
