package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// Create a channel
	myChannel := make(chan string, 5) // create a channel with buffer size of 2
	wg.Add(2)
	// Start a goroutine to send messages on the channel
	go func(myChannel chan<- string) {
		for i := 0; i < 3; i++ {
			myChannel <- fmt.Sprintf("Message %d", i)
		}
		close(myChannel) // Close the channel after sending all messages
		wg.Done()
	}(myChannel)

	// Start a goroutine to receive messages from the channel
	go func(myChannel <-chan string) {
		for msg := range myChannel {
			fmt.Println(msg)
		}
		wg.Done()
	}(myChannel)

	wg.Wait()
	fmt.Println("Done")
}
