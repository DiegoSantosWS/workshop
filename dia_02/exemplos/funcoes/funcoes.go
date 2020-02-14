package main

import (
	"errors"
	"fmt"
)

type myFunc = func(l, b int) int

type area struct {
	Largura int
	Altura  int
}

func (r *area) CalculaArea() int {
	res := r.Largura * r.Altura
	return res
}

func (r area) CalculaPerimetro() int {
	res := 2 * r.Largura * 2 * r.Altura
	return res
}

func main() {
	soma(func(l, b int) int {
		return l + b
	})

	fmt.Println("Resultado é:", multiplica(3, 8))

	fn := exemploAnonimo()
	fmt.Println("Resultado é:", fn())
	fmt.Println("Resultado é:", fn())
	fmt.Println("Resultado é:", fn())

	fmt.Println("Nome é:", exemploNomeado("Marcela"))
	fmt.Println("Nome é:", exemploNomeado("Diego"))
	fmt.Println("Nome é:", exemploNomeado("Francisco"))

	fmt.Println("Resultado é:", exemploVariadico(1, 2))
	fmt.Println("Resultado é:", exemploVariadico(2, 3))
	fmt.Println("Resultado é:", exemploVariadico(3, 4))

	tot, err := exemploVariadicoWithErr(1, 2)
	if err != nil {
		return
	}
	fmt.Println("Resultado é:", tot)

	tot2, err := exemploVariadicoWithErr(2, 3)
	if err != nil {
		return
	}
	fmt.Println("Resultado é:", tot2)

	tot3, err := exemploVariadicoWithErr(3, 4)
	checkErr(err)
	fmt.Println("Resultado é:", tot3)

	a := area{Largura: 10, Altura: 5}
	resultArea := a.CalculaArea()
	fmt.Println("area: ", resultArea)
	perim := &a //repassando os valores
	resultPerim := perim.CalculaPerimetro()
	fmt.Println("perim: ", resultPerim)
}

func soma(fn myFunc) {
	res := fn(1, 3)
	fmt.Println(res)
}

func multiplica(a, b int) int {
	return (a * b)
}

func exemploAnonimo() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func exemploNomeado(str string) (nome string) {
	nome = str
	return
}

func exemploVariadico(numeros ...int) (total int) {
	total = 0
	for _, n := range numeros {
		total += n
	}
	return
}

func exemploVariadicoWithErr(numeros ...int) (total int, err error) {
	total = 0

	for _, n := range numeros {
		total += n
	}

	if total == 0 {
		err = errors.New("O resoltado não pode ser zero")
		return
	}
	return
}

func checkErr(err error) {
	if err != nil {
		return
	}
}
