package goarea

import "math"

// Pi e uma proporcao numerica definida pela relacao entre o perimetro de uma circunfrencia e o seu diametro
const Pi = 3.1416

// Cir e responsavel por carcular area da circunfrencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Retangulo e responsavel por carcular area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Nao e visivel
func _TrianguloEq(base, altura float64) float64 {
	return Rect(base, altura) / 2
}
