package main

import "fmt"

type myFunc = func(l, b int) int

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
