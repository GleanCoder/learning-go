package main

import "fmt"

func main() {
	fmt.Println("Hello Go!")
}

// whenever we run a go program it looks for `package main` which is main package.
// second it looks for `func main()` which is the main function of go programming and also the entry point of the program.

// 👇 Detailed Explanation: 

/*

`Package main`:
- Every Go file belongs to a package.
- package main is special — it tells Go this is an executable program, not a library.
- If you named it something else (like package utils), 
  Go would treat it as a package meant to be imported by other code, and it wouldn't produce a runnable binary.

- summary: package main — required so go build knows to produce an executable instead of a library.
*/

/*

`import "fmt"`:
- This pulls in the fmt package from Go's standard library.
- fmt stands for "format" and contains functions for formatted I/O — printing to the console, reading input, formatting strings, etc.
- You need to import any package whose functions you want to use; Go doesn't have global built-in print functions like some languages do.

- summary: import "fmt" — required because Println doesn't exist unless you import the package that defines it; Go will actually throw a compile error if you import something and don't use it, or use something without importing it.
*/


/*

`func main() {}`:
- This declares the main function. Just like package main is special, the main function inside a main package is special too — it's the entry point of the program.
- When you run the compiled binary, execution starts here from main func.
- If you don't have a func main() in package main, Go won't know where to start and will refuse to build an executable.

- func main() — required as the entry point; without it, go run or go build will fail with "runtime.main_main·f: function main is undeclared in the main package."
*/

/*

`fmt.Println("Hello Go!")`:
- This calls the Println function from the fmt package.
- Println prints its arguments to standard output (your terminal) and adds a newline at the end.
- The syntax fmt.Println means "the Println function that belongs to the fmt package" — Go uses this package.Function syntax to access anything exported (public) from an imported package.

- fmt.Println(...) — this is the actual work being done; without it, the program would compile and run but produce no output.
*/

/*
`How to run Go Program?`

- Go is very strict about unused imports and unused variables — both are compile errors, not warnings.

- we can run go program by using go build filename.go => it will create a executable file and then you can run that file by using ./filename
- or we can run by using go run filename.go

*/