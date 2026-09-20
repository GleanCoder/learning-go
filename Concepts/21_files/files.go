package main

import (
	"bufio"
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

	// file, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()                // close the file after reading the data asap
	// bufferStorage := make([]byte, 12) // buffer is the temporary space created inside the memory to store the data

	// d, err := file.Read(bufferStorage) // read method is used to read the data from the file and store it into the buffer

	// if err != nil {
	// 	panic(err)
	// }
	// for i := 0; i < len(bufferStorage); i++ {
	// 	fmt.Println("Buffer Data:", d, "=", string(bufferStorage[i]))
	// }

	// way 2: Read data using ReadFile method

	// data, err := os.ReadFile("example.txt") // this method is used to read the data from the file and store it into the variable
	// // ReadFile method is not suitable for large files because it reads the entire file into memory, which can lead to high memory usage and potential performance issues. For large files, it's better to read the file in chunks or use a buffered reader.
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	// we can do file streaming how we used to do in nodejs,

	// lets play with folders

	// folder, err := os.Open("../") // open helps to load an existing folder into memory. here we are opening the current folder using "." dot notation

	// if err != nil {
	// 	panic(err)
	// }
	// defer folder.Close() // close the folder after reading the data asap

	// folderInfo, err := folder.ReadDir(-1) // this returns use the information about the folder

	// for _, file := range folderInfo {

	// 	fmt.Println("files:", file) // return folder information
	// }

	// lets see how we can create a new file and write data into it

	// file, err := os.Create("example_two.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer file.Close()

	// file.WriteString("Hey Go")
	// file.WriteString(", you're a great language")

	// whenever we work with files, we work with the binary data, so we can also write the binary data into the file

	// we will add it using byte of slice.

	// byteData :=
	// 	[]byte("Golang Developers")

	// file, err := os.Create("example_three.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// file.Write(byteData)

	// let's see how can we transfer the data from one file to another file using buffer io package.

	sourceFile, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	destFile, err := os.Create("example_four.txt")

	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {
		byt, err := reader.ReadByte()

		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}

			break
		}

		erro := writer.WriteByte(byt)

		if erro != nil {
			panic(erro)
		}

	}
	writer.Flush()

	fmt.Println("Writing new file succesfully done!")

	removeError := os.Remove("example_demo.txt")

	if removeError != nil {
		panic(removeError)
	}

}
