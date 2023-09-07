package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Crea dos canales, uno para "Hola" y otro para "Mundo"
	holaCh := make(chan struct{})
	mundoCh := make(chan struct{})

	// Incrementa el contador del WaitGroup para las dos goroutines
	wg.Add(2)

	// Goroutine para imprimir "Hola"
	go func() {
		defer wg.Done()
		<-holaCh // Espera hasta que se reciba la señal "Hola"
		fmt.Print("Hola ")
		close(mundoCh) // Envía una señal para imprimir "Mundo"
	}()

	// Goroutine para imprimir "Mundo"
	go func() {
		defer wg.Done()
		<-mundoCh // Espera hasta que se reciba la señal "Mundo"
		fmt.Println("Mundo")
		close(holaCh) // Envía una señal para imprimir "Hola"
	}()

	// Inicia la secuencia imprimiendo "Hola" en primer lugar
	close(holaCh)

	// Espera a que ambas goroutines terminen
	wg.Wait()
}
