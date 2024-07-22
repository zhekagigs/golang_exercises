package main

import (
	"fmt"
	"time"
)

const (
	INFO  = "INFO"
	WARN  = "WARNING"
	ERROR = "ERROR"
)

type LogEntry struct {
	msg        string
	timestampt time.Time
	level      string
}

func (le LogEntry) String() string {
	return fmt.Sprintf("[%s] %s - %s", le.level, le.timestampt.UTC(), le.msg)
}

var logCh = make(chan LogEntry, 50)
var doneCh = make(chan struct{})

func main() {
	go logger()
	logCh <- LogEntry{"Application started", time.Now(), INFO}
	time.Sleep(1 * time.Second)
	logCh <- LogEntry{"Error happened", time.Now(), ERROR}
	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{}
}

func logger() {
	for {
		select{
		case entry := <-logCh:
			fmt.Println(entry)
		case <- doneCh:
			break
	}
}
}