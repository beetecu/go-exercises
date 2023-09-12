package main

import (
	"fmt"
	"sync"
	"time"
)

type Box struct {
	mutex   sync.RWMutex
	data    map[string]string
	version int
}

func main() {
	dataset := &Box{
		data: make(map[string]string),
	}

	var wg sync.WaitGroup
	numReaders := 3
	numWriters := 2

	// Reading goroutine
	for i := 1; i <= numReaders; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				dataset.Read(id)
			}
		}(i)
	}

	// Writing goroutine
	for i := 1; i <= numWriters; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				dataset.Write(id)
			}
		}(i)
	}

	wg.Wait()
}

func (d *Box) Read(readerID int) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	fmt.Printf("Reader %d is reading data (version %d): %v\n", readerID, d.version, d.data)
	time.Sleep(time.Millisecond * 500)
}

func (d *Box) Write(writerID int) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.version++
	d.data["key"] = fmt.Sprintf("Value%d", writerID)
	fmt.Printf("Writer %d is writing data (version %d)\n", writerID, d.version)
	time.Sleep(time.Millisecond * 1000)
}
