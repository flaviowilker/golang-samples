package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	ch := make(chan int)

	go receive(ch)
	go send(ch)

	wg.Wait()
	fmt.Printf("%T\n", ch)
}

func send(c chan<- int) {
	defer wg.Done()
	c <- 10
}

func receive(c <-chan int) {
	defer wg.Done()
	fmt.Println(<-c)
}
