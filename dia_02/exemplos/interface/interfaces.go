package main

import "fmt"

// Geo interface base para figuras geométricas
type Geo interface {
	Area() float64
}

// Retangulo representa um retângulo
type Retangulo struct {
	Largura float64
	Altura  float64
}

// Area calcula a are de um retângulo
func (r *Retangulo) Area() float64 {
	res := r.Largura * r.Altura
	return res
}

// Triangulo representa um triângulo
type Triangulo struct {
	Base   float64
	Altura float64
}

// Area calcula a are de um triângulo
func (t *Triangulo) Area() float64 {
	res := (t.Base * t.Altura) / 2
	return res
}

func imprimeArea(g Geo) {
	fmt.Println(g)
	fmt.Println(fmt.Sprintf("Área      : %0.2f", g.Area()))
}

func main() {
	r := Retangulo{
		Altura:  10,
		Largura: 5,
	}

	t := Triangulo{
		Base:   10,
		Altura: 5,
	}

	imprimeArea(&r)
	imprimeArea(&t)
}
