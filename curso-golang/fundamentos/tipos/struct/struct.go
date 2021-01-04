package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

//receiver para manipular a struct
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Lapis",
		preco:    1.80,
		desconto: 0.50,
	}
	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{"Caderno", 130.50, 0.90}

	fmt.Println(produto2, produto2.precoComDesconto())
}
