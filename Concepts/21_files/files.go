package main

import (
	"fmt"
	"os"
)

func main() {

	// here we will practice some operation related  to files, for that we have to need OS package

	// file, err := os.Open("example.txt") // open helps to load an existing file into memory.

	// if err != nil {
	// 	// log panic or handle the error

	// 	panic(err)
	// }

	// fileInfo, err := file.Stat() // this returns use the information about the file

	// fmt.Println(fileInfo.Name())    // return file name
	// fmt.Println(fileInfo.Size())    // return size of the file
	// fmt.Println(fileInfo.IsDir())   // is it a directory ?
	// fmt.Println(fileInfo.Mode())    // return file modification permissions
	// fmt.Println(fileInfo.ModTime()) // last modified time

	// Reading the file data
	// way 1: Read data using buffer storage

	file, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()                // close the file after reading the data asap
	bufferStorage := make([]byte, 12) // buffer is the temporary space created inside the memory to store the data

	d, err := file.Read(bufferStorage) // read method is used to read the data from the file and store it into the buffer

	if err != nil {
		panic(err)
	}
	for i := 0; i < len(bufferStorage); i++ {
		fmt.Println("Buffer Data:", d, "=", string(bufferStorage[i]))
	}

}
