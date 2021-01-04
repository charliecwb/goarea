package main

import (
	"fmt"

	"github.com/cod3rcursos/area/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

//multiplexar - misturar msgs num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(html.Titulo("https://www.google.com", "https://youtube.com"),
		html.Titulo("https://www.cod3r.com.br", "https://amazon.com"))
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
