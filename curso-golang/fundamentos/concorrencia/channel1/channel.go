package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 //enviando dados para o canal
	//<- ch ler valor
	ch <- 2
	fmt.Println(<-ch)
}
