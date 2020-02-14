package main

import "fmt"

func main() {
	exemploIF1()
	exemploIF2()
	exemploIF3()
}

func exemploIF1() {
	var a, b = 2, 4
	c := (a + b)
	if c == 6 {
		fmt.Println("Sua soma está correta.")
		return
	}
	//uma forma d fazer se não
	fmt.Println("Sua soma está errada.")
}

func exemploIF2() {
	var a, b = 2, 4
	c := (a + b)
	if c == 7 {
		fmt.Println("Sua soma está correta.")
	} else {
		//uma forma de fazer o else
		fmt.Println("Sua soma está errada.")
	}
}

func exemploIF3() {
	if 2%2 == 0 {
		fmt.Println("É par.")
	} else {
		fmt.Println("É impar.")
	}

	if num := 2; num < 0 {
		fmt.Println(num, "É negativo.")
	} else if num < 10 {
		fmt.Println(num, "Tem um dígito.")
	} else {
		fmt.Println(num, "Tem vários dígitos.")
	}
}
