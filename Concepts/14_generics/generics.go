package main

import "fmt"

func main() {
	/*
	   - Generics let you write one function or type that works with multiple data types,
	     while still keeping type safety.

	   - That means one function works with any type (int, string, float, etc.), instead of writing a separate function for each type.

	   - Let's take an example:
	   Generics = salt that works with any dish.
	   Salt goes into rice, soup, salad, anything — same salt, different foods.
	   You don't need "rice salt," "soup salt," "salad salt" separately.
	*/

	var nums = []int{5, 6, 7, 10, 1}
	printSlice(nums)
	// printSlice(nums)

	var languages = []string{"Golang", "TypeScript", "Java", "Python"}
	printSlice(languages)
	// printStringSlice(languages)
}

// func printSlice(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

/*
the above function is only printing slice but due to different data types we have to createn different functions.
to tackle this we can use  generics in function.
see the below functions with generics
func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}
	we can use interface{}- empty interface which means also any,
	but it's not good to use any and also we can use any char instead of T but its recomended to use T
- we can define multiple type using | between them, refer below code
*/

func printSlice[T int | string](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}
