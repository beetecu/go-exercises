package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create the channel
	channelProd := make(chan int)
	channelConsumer := make(chan int)

	// Goroutines controller
	var wg sync.WaitGroup

	// Initialize producer
	wg.Add(1)
	go producer(channelProd, &wg, channelConsumer)

	// Initialize consumer
	wg.Add(1)
	go consumer(channelProd, &wg, channelConsumer)

	// Close the channel after both have finished
	close(channelProd)

	// Wait for both producer and consumer to finish
	//wg.Wait()

	// Wait for the done signals
	//<-done
	//<-done

	go func() {
		//defer close(done)
		wg.Wait()
		close(channelConsumer)
	}()

	//close(done)
}

func producer(channel chan<- int, wg *sync.WaitGroup, channelOut chan<- int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Println("We produce the number:", i)
		channelOut <- i // Send the number to the channel
	}
}

func consumer(channel chan<- int, wg *sync.WaitGroup, channelOut chan<- int) {
	defer wg.Done()
	//for num := range channel {
	//	fmt.Println("We have received from channel the number:", num)
	//	}
	fmt.Println("We have received from channel the number:", 1)

	//done <- true // Signal that consumer has finished
}
