package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Number of goroutines required
	numGoroutines := 5

	// Declare
	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Wait the subroutines end
	wg.Wait()

	fmt.Println("Well done. Ending")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Goroutine %d: Time to work (fake work)\n", id)

	// Simulating processing time
	time.Sleep(time.Second * 2)

	fmt.Printf("Goroutine %d: Work ended\n", id)
}
