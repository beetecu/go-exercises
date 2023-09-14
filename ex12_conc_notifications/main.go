package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	numGoroutines := 115

	// Create a channel for notification
	notificationChannel := make(chan struct{})

	// Use a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Create multiple goroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// Simulate some work
			doWork(id)
			fmt.Printf("Goroutine %d has finished\n", id)
			// Notify that this goroutine has finished
			notificationChannel <- struct{}{}
		}(i)
	}

	// Create a separate goroutine to listen for notifications
	go func() {
		wg.Wait()                  // Wait for all goroutines to finish
		close(notificationChannel) // Close the notification channel
	}()

	// Wait for all goroutines to finish and the notification channel to be closed
	<-notificationChannel
	fmt.Println("All goroutines have finished")
}

func doWork(id int) {
	// Simulate some work
	for i := 0; i < 1e8; i++ {
	}
	time.Sleep(time.Millisecond * time.Duration(1000))
}
