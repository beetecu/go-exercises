package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d: Procesing task: %d\n", id, task)
		// Simulating task (replace for a real function implementing the task)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d: Task %d done\n", id, task)
		results <- task
	}
}

func main() {
	numWorkers := 3
	numTasks := 5

	// Create the channels
	tasks := make(chan int, numTasks)
	results := make(chan int, numTasks)

	// goroutines controller
	var wg sync.WaitGroup

	// Create the workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	// Task to the pool
	for i := 1; i <= numTasks; i++ {
		tasks <- i
	}

	// Close tasks channel
	close(tasks)

	// Wait for the all tasks end
	wg.Wait()

	// Close result channel
	close(results)

	// Show taks result
	for result := range results {
		fmt.Printf("Recepted result: %d\n", result)
	}
}
