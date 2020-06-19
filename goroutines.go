package main

import (
	"fmt"
	"sync"
	"time"
)

func fibonacci(num int64) int64 {
	if num < 0 {
		print("Número menor que zero")
		return 0
	} else if num == 0 || num == 1 {
		return 1
	} else {
		return fibonacci(num-1) + fibonacci(num-2)
	}
}

func printFibonacci(index int) {
	defer wg.Done()

	//mude o tamanho desse for para o segundo tamanho do grupo
	for i := 0; i < 10; i++ {
		fmt.Println(index, " - ", i, " - ", fibonacci(35))
	}
    //fmt.Println(index, " - ", 0, " - ", fibonacci(35))
}

var wg sync.WaitGroup

func main() {

	dataAtual := time.Now()
	fmt.Println("Data atual ", dataAtual.String())

	//mude o tamanho desse for para o primeiro tamanho do grupo
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go printFibonacci(i)
	}

	wg.Wait()

	novaDataAtual := time.Now()

	fmt.Println("Tempo em nano segundos (truncado): ", novaDataAtual.Unix()-dataAtual.Unix())
	fmt.Println("Data atual ", novaDataAtual.String())
}

//formato que sera impresso: <num_grupo_1> - <num_grupo_2> - <resultado_fibonacci>
