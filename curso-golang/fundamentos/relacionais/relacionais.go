package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("String:", "" == "")
	fmt.Println("Diferente:", "" != "1")
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println(d1 == d2)

	type Pessoa struct {
		Nome  string
		Idade int
	}

	p1 := Pessoa{"Luana", 40}
	p2 := Pessoa{"Luana", 40}

	fmt.Println(p1 == p2)
}
