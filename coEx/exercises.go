package main
import (
	"time"
	"fmt"
	"context"
)

// Exercise 2: Implement log levels with separate channels
func exerciseLogLevels() {
	infoCh := make(chan LogEntry, 50)
	warnCh := make(chan LogEntry, 50)
	errorCh := make(chan LogEntry, 50)
	// doneCh := make(chan struct{})

	// TODO: Implement three logger functions, one for each level
	// Each logger should read from its respective channel and print entries
	// Use select to handle multiple channels in the main logger function

	// Simulate log entries
	go func() {
		for i := 0; i < 10; i++ {
			infoCh <- LogEntry{fmt.Sprintf("Info log %d", i), time.Now(), INFO}
			warnCh <- LogEntry{fmt.Sprintf("Warning log %d", i), time.Now(), WARN}
			errorCh <- LogEntry{fmt.Sprintf("Error log %d", i), time.Now(), ERROR}
			time.Sleep(100 * time.Millisecond)
		}
		close(infoCh)
		close(warnCh)
		close(errorCh)
	}()

	// TODO: Implement main logger function to handle all channels
	// Use select statement to process logs from different channels
	// Exit when all channels are closed
}

// Exercise 3: Implement a timeout for slow log processing
func exerciseTimeout() {
	logCh := make(chan LogEntry, 50)
	// doneCh := make(chan struct{})

	// TODO: Implement a logger function that processes logs with a timeout
	// If processing a log takes more than 200ms, skip it and log a warning
	// Use select with a time.After channel for the timeout

	// Simulate log entries with random processing times
	go func() {
		for i := 0; i < 10; i++ {
			logCh <- LogEntry{fmt.Sprintf("Log entry %d", i), time.Now(), INFO}
			time.Sleep(time.Duration(100+i*50) * time.Millisecond)
		}
		close(logCh)
	}()

	// TODO: Implement main logger function with timeout handling
}

// Exercise 4: Implement graceful shutdown with context
func exerciseGracefulShutdown() {
	logCh := make(chan LogEntry, 50)
	ctx, _ := context.WithCancel(context.Background())
	
	// TODO: Implement a logger function that respects context cancellation
	// Use the context to signal when the logger should stop processing logs

	// Simulate log entries
	go func() {
		for i := 0; i < 100; i++ {
			select {
			case logCh <- LogEntry{fmt.Sprintf("Log entry %d", i), time.Now(), INFO}:
			case <-ctx.Done():
				return
			}
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// TODO: Implement main function that starts the logger and simulates a shutdown after 2 seconds
	// Use time.After to trigger the shutdown
	// Ensure all queued logs are processed before shutting down
}