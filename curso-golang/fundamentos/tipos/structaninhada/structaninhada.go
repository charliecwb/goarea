package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) totalizarPedido() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}

	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{1, 2, 10.00},
			item{2, 3, 20.50},
			item{3, 4, 5.80},
		},
	}

	fmt.Printf("Total do pedido: %.2f", pedido.totalizarPedido())
}
