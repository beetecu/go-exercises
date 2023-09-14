# go-exercises
Set of solved exercises to learn go

- ex1_conc_num: Create a function that executes a simple goroutine to print numbers 1 to 10 in order.
- ex3_conc_hello_world: Use multiple goroutines to print "Hello" and "World" concurrently and make sure the output is "Hello World" rather than mixed.
- ex4_conc_msg: Program using channels to send and receive messages between two goroutines
- ex5_conc_prod_consumer: Example of producer/consumer 
- ex6_conc_workers: Create several goroutines that perform a fictitious job and wait for them all to finish before printing a completion message.
- ex7_conc_bank_transactions: Create an application that simulates a bank account shared among several goroutines. Use mutexes to avoid race conditions and ensure secure transactions.
- ex8_conc_select: Create a program that uses the select statement to wait for multiple channels and perform actions based on the first one to complete. Implement a timeout so that if none of the goroutines complete within a certain time, a predetermined action is performed.
- ex9_conc_pool: Implement a pool of worker goroutines that can process tasks concurrently. Create a program that sends tasks to the pool and waits for them all to complete.
- ex10_conc_marathon: Design a program that simulates a race between several goroutines. Use channels to coordinate the start and end of the race and determine which goroutine is the winner.
- ex11_conc_mutex: Create a program that uses read-write mutexes to manage concurrent access to a shared data structure. It allows multiple goroutines to read simultaneously, but blocks access during writes.
- ex12_conc_notifications: Experiment with closed channels and channel selection to implement a notification system. Create a program in which several goroutines notify their completion and another goroutine listens to them and displays a message when they are all finished.
- ex13_conc_log_proc: Create a program that processes in streaming mac OS logs in parallel.