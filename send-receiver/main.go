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
	c <- 10
	wg.Done()
}

func receive(c <-chan int) {
	fmt.Println(<-c)
	wg.Done()
}
