package main

import (
	"fmt"
	"time"
)

func main() {

	exemploswitch1()

	exemploswitch2()

	exemploswitch3()
}

func exemploswitch1() {
	i := 2
	switch i {
	case 1:
		fmt.Println("Valor de ", i, " por extenso é: um")
	case 2:
		fmt.Println("Valor de ", i, " por extenso é: dois")
	case 3:
		fmt.Println("Valor de ", i, " por extenso é: três")
	}
}

func exemploswitch2() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("É fim de semana.")
	default:
		fmt.Println("É dia de semana.")
	}
}

func exemploswitch3() {
	j := 3
	switch {
	case 1 == j:
		fmt.Println("Valor por extenso é: um")
	case 2 == j:
		fmt.Println("Valor por extenso é: dois")
	default:
		fmt.Println("Valor não encontrado.")
	}
}
