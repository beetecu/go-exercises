package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create the string channel
	canal := make(chan string)

	// Wait for the subrutines ending
	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine to send the message
	go func() {
		defer wg.Done()
		mensaje := "Hi from subrutine 1"
		canal <- mensaje // Envía el mensaje al canal
	}()

	// Goroutine to receive and print the message
	go func() {
		defer wg.Done()
		mensajeRecibido := <-canal // message catcher
		fmt.Println("The message is here:", mensajeRecibido)
	}()

	// Wait the for the both subroutines end
	wg.Wait()

	// Close the channel
	close(canal)
}
