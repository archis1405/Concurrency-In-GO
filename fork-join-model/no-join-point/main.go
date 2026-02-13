package main

import (
	"fmt"
	"time"
)

func main() {
	go task1()                         // we created a forkpoint here
	time.Sleep(100 * time.Millisecond) // relatively fast task as it takes only 100ms to complete
	fmt.Println("Main function completed, done with the main task")
	// join point is missing here
	/*
		    Output without any join point:
				❯ go run main.go
				  Main function completed, done with the main task
	*/
}

func task1() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Task 1 completed")
}
