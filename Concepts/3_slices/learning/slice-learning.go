package main

import (
	"fmt"
	"slices"
)

func main() {
	/*
		- slices are a key data type in Go, giving a more powerful interface to sequences than arrays.
			- means slices are dynamically-sized arrays. here we don't need to specify the size of the array, we can just use a slice and it will grow as needed.
		- slices are a reference type, meaning that when you assign a slice to another slice, they both point to the same underlying array. changes made to one slice will affect the other.
		- if you don't know the size of the array, you can use slices instead of arrays.
	*/

	/*
		# How to declare a slice
		- you can declare a slice using the following syntax:
			var sliceName []type
		- for example, to declare a slice of integers, you would write:
			var numbers []int
		- you can also use the shorthand syntax to declare and initialize a slice:
			numbers := []int{1, 2, 3, 4, 5}
	*/

	// uninitialized slice is nil and has a length and capacity of 0
	var numbers []int

	fmt.Println(len(numbers), cap(numbers)) // 0 0
	fmt.Println(numbers == nil) // true


	// if we don't want our slice to be nil, we can initialize it with an empty slice literal using the following syntax:
	numbers = []int{}
	fmt.Println(len(numbers), cap(numbers)) // 0 0
	fmt.Println(numbers == nil) // false

	// we can also use the make function to create a slice with a specific length and capacity:
	// if we don't specify the capacity, it will be set to the length of the slice. and also if we don't specify the length, it will be set to 0.
	numbers = make([]int, 5, 10) // length 5, capacity 10
	fmt.Println(len(numbers), cap(numbers))

	// we can also use the append function to add elements to a slice: append  will add elements to the end of the slice and return a new slice with the added elements. if the capacity of the slice is exceeded, a new underlying array will be allocated and the elements will be copied to the new array.
	numbers = append(numbers, 1, 2, 3)
	fmt.Println(len(numbers), cap(numbers))

	// we can also use the copy function to copy elements from one slice to another:
	numbers2 := make([]int, len(numbers))
	// if we don't specify the length of the destination slice, it will be set to the length of the source slice. and also if the destination slice is smaller than the source slice, only the elements that fit in the destination slice will be copied.
	// if we don't specify the capacity of the destination slice, it will be set to the length of the source slice. and also if the destination slice is smaller than the source slice, only the elements that fit in the destination slice will be copied.
	// if destination slice size is 0, then nothing will be copied and the destination slice will remain empty.
	copy(numbers2, numbers)
	fmt.Println(numbers2)

	// we can also use the slicing operator to create a new slice from an existing slice:
	// here 3 index will be excluded, so numbers3 will contain elements at index 1 and 2 of numbers.
	// array_name[from_index:to_index] - from_index is inclusive, to_index is exclusive
	// if we don't specify the from_index, it will be set to 0. and if we don't specify the to_index, it will be set to the length of the slice.
    // if within array if we want elements from a specific range, we can use the slicing operator to create a new slice from an existing slice. for example, if we want elements from index 1 to 3 (excluding 3), we can use the following syntax:
	numbers3 := numbers[1:3]
	fmt.Println(numbers3)

	// initially we have 0 values as per specified length of 5,  because we used make to create the slice with a length of 5. so the first 5 elements are 0, and then we appended 3 more elements to the slice, so the total length is now 8.
	fmt.Println(numbers) // [0 0 0 0 0 1 2 3]


	// we can  fill those 0 values using indexing, for example:
	for  i := 0; i < 5; i++ {
		numbers[i] = i + 10
	}
	fmt.Println(numbers) 


	// normally we keep the length 0 in make function, so an empty slice is created, and then we can append elements to it as needed. this is a common pattern in Go, as it allows us to create a slice with an initial capacity and then grow it as needed.
	numbers4 := make([]int, 0, 10)
	numbers4 = append(numbers4, 1, 2, 3)
	fmt.Println(len(numbers4), cap(numbers4))


	// we have slices package in Go, which provides functions for working with slices. for example, we can use the Equal function to check if two slices are equal, the Clone function to create a copy of a slice, and the Sort function to sort a slice. we can also use the Index function to find the index of an element in a slice, and the Contains function to check if a slice contains an element.

	fmt.Println("Is numbers and numbers2 equal?", slices.Equal(numbers, numbers2))

	// we can make 2D slices also.

	var num= [][]int{{2,3,5},{6,8,10}}
	fmt.Println(num)

	

}