package goarea

import "math"

//Pi é uma const
const Pi = 3.1416

//Circ calcula área de circulo
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect calcula área de retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

//Não é visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
