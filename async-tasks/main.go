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
	done := make(chan struct{})
	now := time.Now()
	fmt.Println("Starting tasks...", time.Since(now))
	go task1(done)
	go task2(done)
	go task3(done)
	go task4(done)

	// time.Sleep(2000 * time.Millisecond) --> this is a temporary solution
	// Why this sleep is needed
	// Because the main function will exit before the goroutines get a chance to complete their tasks
	<-done
	<-done
	<-done
	<-done
	fmt.Println("All tasks completed!", time.Since(now))
}

func task1(done chan struct{}) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task 1 completed")
	done <- struct{}{}
}

func task2(done chan struct{}) {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Task 2 completed")
	done <- struct{}{}
}

func task3(done chan struct{}) {
	fmt.Println("Task 3 completed")
	done <- struct{}{}
}

func task4(done chan struct{}) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task 4 completed")
	done <- struct{}{}
}

/*
	Here the order of the tasks completion is not guaranteed as they are running concurrently
	Task3 (no sleep time) --> Task1 (100ms) --> Task2 (200ms) --> Task4 (400ms)
	The total time taken is the maximum time taken by any of the tasks which is 400ms
	the order of task is not always fixed
*/

//To handle this we will use something known as Fork and Join model
// GO uses fork and join model to handle concurrent tasks
