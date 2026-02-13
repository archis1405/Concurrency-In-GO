package main

import (
	"fmt"
	"time"
)

/* Here each task is waiting for the previous one to get compleated before starting
so the total time taken is the sum of all the tasks time which is 100+200+1+400 = 701ms */

func main() {
	now := time.Now()
	fmt.Println("Starting tasks...", time.Since(now))
	task1()
	task2()
	task3()
	task4()
	fmt.Println("All tasks completed!", time.Since(now))
}

func task1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task 1 completed")
}

func task2() {
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Task 2 completed")
}

func task3() {
	time.Sleep(time.Millisecond)
	fmt.Println("Task 3 completed")
}

func task4() {
	time.Sleep(400 * time.Millisecond)
	fmt.Println("Task 4 completed")
}

/*
Here the order of the tasks completion is guaranteed as they are running synchronously
Task1 (100ms) --> Task2 (200ms) --> Task3 (1ms) --> Task4 (400ms)
The total time taken is the sum of all the tasks time which is 100+200+1+400 = 701ms
*/
