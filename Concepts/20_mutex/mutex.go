package main

import (
	"fmt"
	"sync"
)

type post struct {
	like int
	mu   sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer wg.Done()
	p.mu.Lock() // Lock the mutex to ensure exclusive access to the like count. one goroutine can access the critical section at a time.
	p.like += 1
	defer p.mu.Unlock() // Unlock the mutex to allow other goroutines to access the like count. and used defer to ensure that the mutex is always unlocked, even if an error occurs or the function returns early.
}

func main() {

	/*
		mutex:
		- mutex is a mutual exclusion lock.
		- It is used to protect shared resources from concurrent access by multiple goroutines.


		- suppose we have a social media application where multiple users can like a post simultaneously.
		  If we don't use a mutex, the like count may not be updated correctly, leading to inconsistent data.
		  and here we faced a race condition problem, where multiple goroutines are trying to update the like count simultaneously,
		   leading to unexpected results.


	*/

	var wg sync.WaitGroup

	myPost := post{
		like: 0,
	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}
	wg.Wait()
	fmt.Printf("You got %d like on your recent post!", myPost.like)

}
