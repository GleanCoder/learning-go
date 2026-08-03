package main

import "fmt"

var proceedExit int=1;



func main() {
	// var variable_name variable_type/data_type
	var programmingLanguage string = "TypeScript"
	fmt.Println(programmingLanguage)

	programmingLanguage="JS"
	fmt.Println(programmingLanguage);

	//programmingLanguage:="Java" // here it gives me error like you cann't declare a new variable using the name programmingLanguage cause it's already declared using var.
	
	year:=2026
	fmt.Println(year)
	year=2027
	fmt.Println(year)
	//year := 2028 // not allowed

	fmt.Println("proceed exit before initialize inside func:", proceedExit)
	proceedExit=0
	fmt.Println("proceed exit after initialize inside func: :", proceedExit)
	proceedExit:=2
	fmt.Println("proceed exit after new declaration using shorthand operator", proceedExit)
	

	var ExportedVariable="Hello I exported from variable file"
	var exportedVariable= "Hello I want to be exported"
	fmt.Println(ExportedVariable , exportedVariable)
	
}

/*
we have types like

1. string
2. int
3. float64
4. bool


- we can declare a variable using var keyword and also we can use := operator to declare and initialize a  variable in a single line.
- the only difference between var and := is that var can be used to declare a variable without initializing it, while := requires initialization at the time of declaration.
- another difference is that var can be used to declare a variable at the package level, while := can only be used within a function.
- constants are declared using the const keyword, and their values cannot be changed after they are set. Constants can be of any data type, including string, int, float64, and bool.
- if we use variable name with the first letter in uppercase, it will be exported and can be accessed from other packages. If the first letter is lowercase, it will be unexported and can only be accessed within the same package.


*/

/*
What came to know while exploring:

- In Go, you cannot call fmt functions (like fmt.Println or fmt.Printf) directly at the package level (outside of a function body) because Go does not allow executable statements at the package scope.
- At the package level, you can only declare constants, variables, types, and functions. You cannot write procedural code that executes an action.
*/