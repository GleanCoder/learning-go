package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		- Goroutines are lightweight threads managed by the Go runtime. They allow concurrent execution of functions, enabling efficient multitasking within a Go program.
		- Goroutines are created using the `go` keyword followed by a function call. When a goroutine is started, it runs concurrently with other goroutines in the same program.
		- Goroutines are multiplexed onto a smaller number of operating system threads, which makes them more memory-efficient compared to traditional threads.
		- Communication between goroutines is often done using channels, which provide a safe way to send and receive data between concurrent functions.
		- The Go runtime scheduler manages the execution of goroutines, allowing them to be paused and resumed as needed, ensuring efficient use of system resources.

		- what is go runtime scheduler?
		- runtime scheduler is a component of the Go programming language that manages the execution of goroutines. It is responsible for scheduling and coordinating the execution of multiple goroutines on available operating system threads. The scheduler ensures that goroutines are executed efficiently, allowing for concurrent execution while minimizing resource usage.

		- How it works:
		1. The Go runtime maintains a pool of operating system threads, which are used to execute goroutines.
		2. When a goroutine is created, it is added to a queue of runnable goroutines.
		3. The scheduler selects goroutines from the queue and assigns them to available threads for execution.
		4. If a goroutine performs a blocking operation (e.g., waiting for I/O), the scheduler can pause that goroutine and schedule another one to run on the same thread.
		5. The scheduler uses a work-stealing algorithm to balance the load across threads, ensuring that all available resources are utilized effectively.
	*/

	for i := 0; i <= 10; i++ {
		go taskExecutor(i)
	}

	// our main function will exit before the goroutines have a chance to complete their execution. To prevent this, we can use a simple mechanism to wait for all goroutines to finish before exiting the main function.

	time.Sleep(2 * time.Second) // Wait for 2 seconds to allow goroutines to complete

}

func taskExecutor(taskNo int) {
	fmt.Println("Executing Task No.: ", taskNo)
}
