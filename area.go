package area

import "math"

//Pi é uma proporçao numerica definifa pela relação entre o perímetro de uma circunferência
//e seu diâmetro
const Pi = 3.1416

//Circ é responsavel por calcular a area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect é responsavel por calcular a area de um retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
