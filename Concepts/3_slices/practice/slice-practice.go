package main

import "fmt"

func main() {
	// slice is a dynamic array which can grow and shrink in size. not limited like array.
	// we can declare slice in two ways
	// 1. using slice literal
	var num = []int{}
	fmt.Println("Length of num:", len(num))
	fmt.Println("Capacity of num:", cap(num))


// 2. using make function
	
var num1= make([] int, 4)
fmt.Println(len(num1))
fmt.Println(cap(num1))
fmt.Println(num1)
// num1[0]=3
num2:=append(num1,7)
num2= append(num2, 8)
num2=append(num2, 4)
fmt.Println(len(num2))
fmt.Println(cap(num2))
fmt.Println(num2)

num3:=num2
num3[5]=67
fmt.Println(num2)
/*
- if we want to specify some initial value of length or to create a slice with non-zero length then we can specify those using make method.

- variable_name=make(array_type,length,capacity)
- if we don't specify the capacity then the capacity will be same as length
-  the moment when no. of element exceed the limit of capacity, it will double the capacity.
- if we did num1.append(value) => it will be  [0,0,0,value]
- slice use reference of the orignal array underlaying, if you change in one then it will reflects in other.


*/


}