package main

import (
	"fmt"
	"maps"
)

func main() {
	/*
	   - How we are using hash and object in js and dictionary in python, similarly we can use map in golang

	   - to create map we will use make() in golang

	   Syntax:

	   make(map[type of element] type of value) => make(map[string] int)

	*/

	var personalDetails = make(map[string]int)

	// to set values: map_variable_name["element_name"]=value

	personalDetails["age"] = 18

	// to get value: map_variable_name["element_name"]

	fmt.Println(personalDetails["age"])

	/* 
	- suppose an element doesn't exist and we try to get that then it will give us the zero value as per the type 
	- int => 0
	- string => empty string
	*/
	fmt.Println(personalDetails["exist"]) // 0

	personalDetails["year"]=2002
	personalDetails["date"]=11


	// we can check map length using len method
fmt.Println(len(personalDetails))


// we can delete an element from map also using delete method: delete(map_name,"element_name")

delete(personalDetails,"date")
fmt.Println(len(personalDetails))
fmt.Println(personalDetails)


// we can also clear the map using clear method

// clear(personalDetails)

// fmt.Println(len(personalDetails))
// fmt.Println(personalDetails)


// we can declare map other than use make method: if you know or have the element in advance then use the below way, otherwise use the make method.

// syntax: variable_name := map[type_of_element] type_of_value {element:value,element2:value2,...}

productList := map[int] string{1:"book",2:"pen"}
fmt.Println(productList)


// suppose sometime we have to check whether an element is present or not in a map and based on that we will do some operations, so how can we do it ?


// we can use any other name than ok but in go it's idiomatic to use ok and also if the element is not present in map then the value will be zero value
 value,ok :=personalDetails["date"]
fmt.Println("Value:",value)
 if !ok {
	fmt.Println("Element is not present!")
 } else{
	fmt.Println("Element is present!")
 }


 // we can also compare two maps in golang similarly how we use slices package in slices, here in map we can use maps package

 var map1 = map[string] int{"date":17,"year":2026}
  var map2 = map[string] int{"date":17,"year":2026}

  fmt.Println("Is maps were equal?", maps.Equal(map1,map2))
 



}