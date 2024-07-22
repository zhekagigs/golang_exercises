package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

/**
This is the main function of the program, which simulates concurrent increments to a shared variable.
It uses two goroutines: one for printing hello messages and another for incrementing a counter.
The WaitGroup is used to wait until all goroutines have finished their tasks.

@main
@package main
@return none
@description Concurrently prints "Hello" 10 times, while also concurrently incrementing a shared variable.
*/
func main(){
	for i := 0; i < 10; i++ {
		wg.Add(2)
		// m.RLock()
		go sayHello()
		// m.Lock()
		go increment()
	}
	wg.Wait()
}


func sayHello(){
	fmt.Printf("Hello %v\n", counter)
	// m.RUnlock()
	wg.Done()
}


func increment(){
	counter++
	// m.Unlock()
	wg.Done()
}
