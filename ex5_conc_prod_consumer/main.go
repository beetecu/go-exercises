package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	messages := make(chan string)

	// Start messageReceiever goroutines
	for x := 1; x <= 4; x++ {
		wg.Add(1)
		go messageReceiever(x, wg, messages)
	}

	// Use a separate goroutine to close the messages channel
	go func() {
		wg.Wait()
		close(messages)
	}()

	// Read from the messages channel
	for msg := range messages {
		fmt.Println(msg)
	}
}

func messageReceiever(count int, wg *sync.WaitGroup, messages chan string) {
	defer wg.Done()
	//time.Sleep(time.Millisecond * time.Duration(1000))
	messages <- fmt.Sprintf("John Wick: %d", count)
}
