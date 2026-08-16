package main

import "fmt"

func main() {
	// range use to iterate over different types of data structures like slice,map and string

	/*
		1. Slice:

	*/

	num := []int{7, 8, 10}

	// either we can iterate this using for loop or we can use range which is more effective

	for idx, val := range num {
		fmt.Println("index:", idx, " =>", val)

	}

	// 2. Map:

	capitalOfStates := map[string]string{"Odisha": "BBSR", "Maharashtra": "Mumbai", "WB": "Kolkata", "Jharkhand": "Ranchi"}

	for key, value := range capitalOfStates {
		fmt.Println(key, ":", value)
	}

	// it will print only element names not value
	for k := range capitalOfStates {
		fmt.Println(k)
	}

	// 3. String:
	str := "Hello"
	for idx, char := range str {
		fmt.Println("index:", idx, " =>", string(char))
	}

	/*
			 here idx is the index of the character in the string and char is the unicode value of the character. We need to convert it to string to print it as a character.
			 - this unicode is runes which is a data type in go that represents a unicode code point. It is an alias for int32. Each character in a string is represented by a rune value. The range loop iterates over the string and returns the index and the rune value of each character.
			 - and if the range comes under 255 then it will be represented as a byte which is an alias for uint8. It is used to represent ASCII characters. The range loop will return the index and the byte value of each character in the string.
			 - if it cross es 255 then it will be represented as a rune which is an alias for int32. It is used to represent Unicode characters. The range loop will return the index and the rune value of each character in the string.
		     - it will take 2 bytes and the index will be incremented by 2 for that character. For example, the character '€' has a unicode value of 8364 which is greater than 255. So it will be represented as a rune and the index will be incremented by 2 for that character.
	*/

}
