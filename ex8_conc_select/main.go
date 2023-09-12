package main

import (
	"fmt"
	"time"
)

func main() {
	// Create two channels
	canal1 := make(chan string)
	canal2 := make(chan string)

	// Goroutine 1: Simulate a task 1
	go func() {
		time.Sleep(5 * time.Second)
		canal1 <- "Task 1 -> done"
	}()

	// Goroutine 2:  Simulate task 2
	go func() {
		time.Sleep(6 * time.Second)
		canal2 <- "Task 2 -> done"
	}()

	// Utilizar la declaración select para esperar el primero que se complete
	select {
	case result := <-canal1:
		fmt.Println(result)
	case result := <-canal2:
		fmt.Println(result)
	case <-time.After(4 * time.Second): // Timeout of 4 seconds
		fmt.Println("Tasks running out of time.")
	}
}
