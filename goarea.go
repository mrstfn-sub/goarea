package goarea

import "math"

// Pi é uma proporção numérica definida pela relação
// entre o perímetro de uma circunferência e o seu diâmetro
const Pi = 3.1416

// Circ é uma função responsável por calcular a área da circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect  calcula área de um retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visível (começa com _)
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}

// comentário teste
