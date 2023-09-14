package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/fsnotify/fsnotify"
)

func main() {
	// Root of the system log file for mac os.
	logFilePath := "/var/log/system.log"

	// Open the log file in reading mode
	file, err := os.Open(logFilePath)
	if err != nil {
		log.Fatalf("Error opening the file: %v", err)
	}
	defer file.Close()

	// Create a channel for log processing.
	lines := make(chan string)

	// goroutines controller
	var wg sync.WaitGroup

	// Number of goroutines
	numGoroutines := 4

	// Multiples goroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go processLogLines(lines, &wg)
	}

	// Watcher for file changes detection
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("Watcher creation error : %v", err)
	}
	defer watcher.Close()

	err = watcher.Add(logFilePath)
	if err != nil {
		log.Fatalf("Error adding file to the watcher : %v", err)
	}

	// Process the log.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines <- line
	}

	// Monitor the logs file and read the need logs
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					// Process the new log(s) becase a file change have been detected
					scanner := bufio.NewScanner(file)
					for scanner.Scan() {
						line := scanner.Text()
						lines <- line
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("Watcher error:", err)
			}
		}
	}()

	// Wait the end of all goroutines
	wg.Wait()
}

// processLogLines procesa líneas de registro en paralelo.
func processLogLines(lines <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for line := range lines {
		// Here -> time to process the log
		fmt.Println("Procesando línea de registro:", line)
	}
}
