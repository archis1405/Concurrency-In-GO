package main

import (
	"fmt"
	"time"
)

/*
Here each task is running concurrently without waiting for the previous one to get compleated
So the total time taken is the maximum time taken by any of the tasks which is 400ms
*/

func main() {
	now := time.Now()
	fmt.Println("Starting tasks...", time.Since(now))
	go task1()
	go task2()
	go task3()
	go task4()
	time.Sleep(time.Second)
	// time.Sleep(2000 * time.Millisecond) --> this is a temporary solution
	// Why this sleep is needed
	// Because the main function will exit before the goroutines get a chance to complete their tasks
	fmt.Println("All tasks completed!", time.Since(now))
}

func task1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task 1 completed")
}

func task2() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Task 2 completed")
}

func task3() {
	fmt.Println("Task 3 completed")
}

func task4() {
	time.Sleep(400 * time.Millisecond)
	fmt.Println("Task 4 completed")
}

/*
	Here the order of the tasks completion is not guaranteed as they are running concurrently
	Task3 (no sleep time) --> Task1 (100ms) --> Task2 (200ms) --> Task4 (400ms)
	The total time taken is the maximum time taken by any of the tasks which is 400ms
	the order of task is not always fixed
*/

//To handle this we will use something known as Fork and Join model
// GO uses fork and join model to handle concurrent tasks
