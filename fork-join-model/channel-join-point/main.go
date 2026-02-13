package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	done := make(chan struct{}) // we are creating a channel to signal the completion of the task1() function

	go func() {
		task1()
		done <- struct{}{} // we are sending an empty struct to the channel to signal that the task1() function is completed
	}()

	<-done // we are waiting for the signal from the channel before proceeding

	fmt.Println("Main function completed, done with the main task")
	fmt.Println("Total time taken:", time.Since(now))
}

func task1() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Task 1 completed")
}

// channel join point is faster than waitgroup join point because it does not have the overhead of managing the waitgroup counter, it simply waits for a signal from the channel, whereas waitgroup has to manage the counter and check for completion of tasks.
