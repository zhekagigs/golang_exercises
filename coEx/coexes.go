package main

import (
	// "context"
	"fmt"
	"sync"
	"time"
)

const (
	INFO  = "INFO"
	WARN  = "WARNING"
	ERROR = "ERROR"
)

type LogEntry struct {
	msg       string
	timestamp time.Time
	level     string
}

func (le LogEntry) String() string {
	return fmt.Sprintf("[%s] %s - %s", le.level, le.timestamp.UTC().Format(time.RFC3339), le.msg)
}

var wg = sync.WaitGroup{}

func worker(logCh <-chan LogEntry, doneCh chan<- struct{}, i int) {
	for msg := range logCh {
		time.Sleep(1 *time.Millisecond)
		fmt.Println(i, msg)
	}
	wg.Done()
	doneCh <- struct{}{}
}

// Exercise 1: Implement a worker pool for logging
func exerciseWorkerPool() {
	logCh := make(chan LogEntry, 100)
	doneCh := make(chan struct{})
	wg.Add(3)
	defer wg.Wait()
	for i := 0; i < 3; i++ {
		go worker(logCh, doneCh, i)
	}

	// TODO: Implement a worker pool with 3 worker goroutines
	// Each worker should read from logCh and print the log entries
	// Use a WaitGroup to ensure all workers finish before the main goroutine exits

	// Simulate log entries
	for i := 0; i < 20; i++ {
		logCh <- LogEntry{fmt.Sprintf("Log entry %d", i), time.Now(), INFO}
	}
	close(logCh)
	// TODO: Wait for all workers to finish
	// close(doneCh)
}


func main() {
	fmt.Println("Exercise 1: Worker Pool")
	exerciseWorkerPool()

	fmt.Println("\nExercise 2: Log Levels")
	exerciseLogLevels()

	fmt.Println("\nExercise 3: Timeout")
	exerciseTimeout()

	fmt.Println("\nExercise 4: Graceful Shutdown")
	exerciseGracefulShutdown()
}
