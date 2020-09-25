package main

import (
	"fmt"
	"sync"
	"time"
)

func fibonacci(num int64) int64 {
	if num < 0 {
		print("The number is lower than zero")
		return 0
	} else if num == 0 || num == 1 {
		return 1
	} else {
		return fibonacci(num-1) + fibonacci(num-2)
	}
}

func printFibonacci(index int) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		fmt.Println(index, " - ", i, " - ", fibonacci(35))
	}
}

var wg sync.WaitGroup

func main() {

	currentDate := time.Now()
	fmt.Println("Current date ", currentDate.String())

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go printFibonacci(i)
	}

	wg.Wait()

	newCurrentDate := time.Now()

	fmt.Println("Time in nano seconds (truncated): ", newCurrentDate.Unix()-currentDate.Unix())
	fmt.Println("Current date ", newCurrentDate.String())
}
