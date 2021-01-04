package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	//inteiro positivo...uint8, uint16, uint32, uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	//inteiro com sinal...int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor maximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a'
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	//numeros reais float32, float64
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo literal de 49.99 é", reflect.TypeOf(49.99))

	//Boolean
	//bo := true

	//string
	//s1 := "Meu nome é Charles"

	//string com multiplas linhas
	s2 := `teste
	de
	quebra`
	fmt.Println(s2)
}
