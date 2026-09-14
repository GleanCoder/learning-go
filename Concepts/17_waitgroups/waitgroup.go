package main

import (
	"fmt"
	"sync"
)

func printTasks(id int, wg *sync.WaitGroup) {
	defer wg.Done() // it will execute after the function returns its result and decrease Add by -1 so it will be 0 and goroutine gets end
	// we have to use defer to tell, execute after the function returns its final value, I mean after execution of the function
	fmt.Println("Doing Task:", id)
}

func main() {
	// Create waitgroup
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1) // we use Add to tell go runtime or waitgroup for how many goroutine it have to wait.
		go printTasks(i, &wg)
	}

	/*
		we know the momemnt main exit, it will not wait for the other goroutine to be execute successfully,
		 for that we have to delay it and we have used delay timer method, but that's not appropriate always
		 - why? because we don't know how much time it gonna take,
		 - for that we have something called waitgroup, lets use that

	*/

	wg.Wait() // it tells goroutine to wait untill all goroutine to be finished which  are being tracked by waitgroup
}
