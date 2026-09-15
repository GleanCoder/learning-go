package main

import "fmt"

// send channel

func greetPeoples(name chan string) {
	fmt.Println(<-name)
}

// Recieve channel:

func calcSum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult
}

func main() {
	// channels are the pipelines that make it possible to send and receive values with the channel operator, <-.
	// basically, channels are a way to communicate between goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.

	// How to create channel

	// channelOne := make(chan string) // unbuffered channel

	// How to send data to channel:
	// channelOne <- "Ping"
	// How to recieve data in channel:

	// <-channelOne

	// we can store it another variable

	// msg := <-channelOne

	// fmt.Println(msg)

	nameChannel := make(chan string)

	go greetPeoples(nameChannel)

	nameChannel <- "Aditya"

	result := make(chan int)
	go calcSum(result, 5, 6)

	res := <-result
	fmt.Println(res)
}

/*
📝 Point to Note:
An unbuffered channel in Go does not have any storage capacity,
so when a goroutine sends a value using ch <- value,
the sender blocks until another goroutine is ready to receive that value using <-ch.
Similarly, the receiver blocks until a sender provides a value.
Therefore, an unbuffered channel works like a direct handoff or handshake between goroutines.
 If the sender and receiver are placed sequentially in the same goroutine,
 the sender blocks first, preventing the receiver from ever being reached,
 which results in a deadlock. To avoid this, sending and receiving should happen concurrently,
  usually by using another goroutine.

*/
