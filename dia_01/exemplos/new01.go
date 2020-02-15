package main

import "fmt"

func main() {
	// epIV, do tipo *int, aponta para uma variável sem nome
	epIV := new(int)
	// eraOuroSith, do tipo *int, também aponta para uma variável sem nome
	eraOuroSith := new(int)
	// "0" zero
	fmt.Println(*eraOuroSith)
	// novo valor para o int sem nome
	*eraOuroSith = *epIV - 5000
	// "-5000"
	fmt.Println(*eraOuroSith)
}
