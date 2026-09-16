package main

import (
	"fmt"
)

/*
- In unbuffered channel, the sender must wait until another goroutine receives concurrently.
- otherwise it will face deadlock.

- In buffered channel we can perform a task without a receiver immediately reading them.
- But, if the buffer capacity get full or reached then the rest whatever will be come next will be blocked.

Buffer channel Syntax: make(chan Type, N) => here N is the capacity, sender will waits for receiever only once the capacity of  buffer is full.
- we can use this to implement Queues Systems / work distribution

*/

// lets build a queue system with an example of build email sender

// we can make type safety more stricter for channel, here below
// emailChan and done channel can do both send and recieve but
// here for our usecase emailChan is recieving while done is sending
// we can define this by using <-

func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("Sending email to", email)
		// time.Sleep(time.Second)
	}
	// <-done // invalid operation: cannot receive from send-only channel

}

func main() {
	emailChan := make(chan string, 100)
	done := make(chan bool)
	// if we know how many channels we have send then we don't need close method explicitly
	// emailChan <- "qa.test@gmail.com"
	// emailChan <- "tue.tester@gmail.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)
	go emailSender(emailChan, done)
	for i := 1; i <= 30; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}
	close(emailChan) // if we are using for loop or range then we have to close the channel, cause if we don't close even after buffer sender get finished it task, reciever stills think there is some value going to be recieved and it waits while there are nothing to send. so it causes deadlock, to prevent this we have to use close method.

	<-done

	// we  can recieve multiple different goroutine channels using a combination of select

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 11
	}()

	go func() {
		chan2 <- "ping"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Value := <-chan1:
			fmt.Println("Value of chan1:", chan1Value)

		case chan2Value := <-chan2:
			fmt.Println("Value of Chan2:", chan2Value)
		}
	}

}
