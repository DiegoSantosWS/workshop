package main

import (
	"fmt"
	"time"
)

func main() {

	exemploFor1()
	exemploFor2()
	exemploFor3()
	defer func() {
		time.Sleep(time.Second * 1)
		exemploFor4()
	}()
}

func exemploFor1() {
	for i := 1; i <= 10; i++ {
		fmt.Println("O número é: ", i)
	}
}

func exemploFor2() {
	i := 5
	for i <= 5 {
		fmt.Println("O número é: ", i)
		i = i + 1
	}
}

func exemploFor3() {
	for {
		fmt.Println("Olá sou o infinito")
		break
	}
}

/*============================================== FOR RANGE =============================*/

func exemploFor4() {
	listaDeCompras := []string{"arroz", "fejão", "melancia", "banana", "maçã", "ovo", "cenora"}
	for k, p := range listaDeCompras {
		retornaNomeVegetal(k, p)
	}
}

func retornaNomeVegetal(key int, str string) {
	switch str {
	case "banana", "cenora":
		fmt.Println("A sessão", key, " e vejetal é:", str)
	default:
		return
	}
}
