package main

import (
	"fmt"
	"time"
)

// Message is a custom struct to be used with channels
type Message struct {
	From    string
	Content string
	Time    time.Time
}

type Post struct {
	Id int
	Content string
}

func main() {
	// 1. Channel of strings
	stringCh := make(chan string)
	
	postageChannel := make(chan Post)
	
	go func (postageChannel chan<- Post)  {
		postageChannel <- Post{12, "Postage"}
	}(postageChannel)
	postageReciver := <-postageChannel
	
	fmt.Println(postageReciver)
	go func() {
		stringCh <- "Hello, Gopher!"
	}()
	greeting := <-stringCh
	fmt.Println("Received:", greeting)

	// 2. Channel of custom struct
	messageCh := make(chan Message)

	go func() {
		messageCh <- Message{
			From:    "Alice",
			Content: "How are you?",
			Time:    time.Now(),
		}
	}()
	msg := <-messageCh
	fmt.Printf("Message from %s: %s (at %v)\n", msg.From, msg.Content, msg.Time)

	// 3. Buffered channel of strings
	bufferedStringCh := make(chan string, 2)
	bufferedStringCh <- "Go"
	bufferedStringCh <- "Channels"
	fmt.Println(<-bufferedStringCh, <-bufferedStringCh)

	// 4. Closing a channel of custom structs
	notifications := make(chan Message, 5)
	go func() {
		for i := 1; i <= 3; i++ {
			notifications <- Message{
				From:    fmt.Sprintf("System-%d", i),
				Content: fmt.Sprintf("Notification %d", i),
				Time:    time.Now(),
			}
		}
		close(notifications)
	}()

	// 5. Ranging over a channel of custom structs
	for notification := range notifications {
		fmt.Printf("From: %s, Content: %s\n", notification.From, notification.Content)
	}

	// 6. Select statement with string channels
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Quick response"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Slow response"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received from ch1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received from ch2:", msg2)
		}
	}

	// 7. Non-blocking channel operations with custom struct
	msgCh := make(chan Message)
	select {
	case msg := <-msgCh:
		fmt.Printf("Received: %+v\n", msg)
	default:
		fmt.Println("No message received")
	}

	// 8. Directional channels with custom struct
	sendOnly := make(chan<- Message)
	receiveOnly := make(<-chan Message)

	go sendMessage(sendOnly)
	go receiveMessage(receiveOnly)

	time.Sleep(1 * time.Second) // Allow goroutines to execute
}

func sendMessage(ch chan<- Message) {
	ch <- Message{
		From:    "Bob",
		Content: "Hello!",
		Time:    time.Now(),
	}
}

func receiveMessage(ch <-chan Message) {
	msg := <-ch
	fmt.Printf("Received message: %+v\n", msg)
}