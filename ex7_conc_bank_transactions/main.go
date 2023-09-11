package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (c *BankAccount) Deposit(amount int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.balance += amount
	fmt.Printf("Added $%d done. Curent balance: $%d\n", amount, c.balance)
}

func (c *BankAccount) Withdraw(amount int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.balance >= amount {
		c.balance -= amount
		fmt.Printf("Withdraw $%d done. Curent balance: $%d\n", amount, c.balance)
	} else {
		fmt.Printf("balance insuficient for withdraw $%d. Curent balance: $%d\n", amount, c.balance)
	}
}

func main() {
	account := BankAccount{balance: 1000}
	var wg sync.WaitGroup

	// Simular múltiples transacciones concurrentes
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			account.Deposit(100)
			account.Withdraw(50)
		}(i)
	}

	// Wait for gorutines end
	wg.Wait()

	fmt.Printf("Final balance: $%d\n", account.balance)
}
