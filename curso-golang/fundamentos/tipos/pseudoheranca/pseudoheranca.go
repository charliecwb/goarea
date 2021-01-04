package main

import "fmt"

type carro struct {
	nome       string
	velocidade int
}

type ferrari struct {
	carro       //campos anonimos
	turboligado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidade = 200
	f.turboligado = true

	fmt.Println(f)
}
