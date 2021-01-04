package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second * 2)
	c <- 1
	fmt.Println("Só depois que foi lido")
}

func main() {
	c := make(chan int) //canal sem buffer

	go rotina(c)

	fmt.Println(<-c)
	fmt.Println("Foi lido")
	fmt.Println(<-c) //deadlock - depois que foi lido uma vez, não lê mais
	fmt.Println("Fim")
}
