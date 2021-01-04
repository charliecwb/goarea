package main

import "fmt"

func main() {
	i := 1

	//go não tem aritmetica de ponteiro
	var p *int = nil

	p = &i //pegando o endereço de i
	*p++
	i++
	fmt.Println(p, *p, i, &i)
}
