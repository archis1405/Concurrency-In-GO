package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()

	// we are using waitgroup to create a join point here, so that the main function will wait for the task1() to complete before it exits
	var waitgroup sync.WaitGroup

	waitgroup.Add(1) // we are adding a task to the waitgroup
	// parameter of Add() is the number of tasks that we are adding to the waitgroup,
	// in this case we are adding one task which is task1()

	go func() {
		defer waitgroup.Done() // we are marking the task as done when it is completed, this will decrement the waitgroup counter by 1
		task1()                // we created a forkpoint here
	}()

	waitgroup.Wait() // this will block the main function until all tasks are done i.e wait until all the tasks are completed
	fmt.Println("Main function completed, done with the main task")
	fmt.Println("Total time taken:", time.Since(now))
}

func task1() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Task 1 completed")
}

// defer waitgroup.Done() is used to ensure that the Done() method is called even if there is an error in the task1() function, this way we can avoid deadlocks in case of errors.
