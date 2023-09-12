package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {

	numGoroutines := 5

	// Channel for the runners
	init := make(chan struct{})

	// Winner results channel
	result := make(chan int)

	// Goroutines controller
	var wg sync.WaitGroup

	// Start signal
	fmt.Println("Start signal!")

	for i := 1; i <= numGoroutines; i++ {
		wg.Add(1)
		go func(runnerId int) {
			defer wg.Done()
			// Wait for the start signal
			<-init
			distance := rand.Intn(100) // Random distance
			fmt.Printf("Goroutine %d: Running %d meters\n", runnerId, distance)
			result <- distance
		}(i)
	}

	close(init)

	// wait for the end of all goroutines
	go func() {
		wg.Wait()
		close(result) // Close the result channel
	}()

	// Getting the winner
	var theWinner int
	var maxDistance int
	var runnersEnd int

	for distance := range result {
		runnersEnd++
		if distance > maxDistance {
			maxDistance = distance
			theWinner = runnersEnd
		}
	}

	fmt.Printf("¡The runner %d is the winner with %d meters!\n", theWinner, maxDistance)
}
