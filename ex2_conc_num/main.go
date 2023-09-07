package main

import (
	"fmt"
	"sync"
)

func imprimirNumeros(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	var wg sync.WaitGroup

	// Increase counter
	wg.Add(1)

	// Launch goroutine
	go imprimirNumeros(&wg)

	// Wait for the gorutine end
	wg.Wait()
}
