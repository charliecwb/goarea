package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type composicao interface {
	esportivo
	luxuoso
}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	fmt.Println("Turbo")
}

func (b bmw7) fazerBaliza() {
	fmt.Println("Baliza")
}

func main() {
	var bmw composicao = bmw7{}
	bmw.ligarTurbo()
	bmw.fazerBaliza()
}
