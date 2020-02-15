package main

import "fmt"

func main() {
	eraOuroSith, epIV := 42, 37

	// ponteiro para eraOuroSith
	p := &eraOuroSith
	// valor de eraOuroSith por meio do ponteiro
	fmt.Printf("Era de Ouro dos Sith - %d anos antes do Ep.IV", *p)
	// atualiza o valor de eraOuroSith por meio do ponteiro
	*p = 5000
	// o novo valor de eraOuroSith
	fmt.Printf("Era de Ouro dos Sith - %d anos antes do Ep.IV | Atualizado", *p)

	// ponteiro para epIV
	p = &epIV
	// divide epIV por meio do ponteiro
	*p = *p / 37
	// o novo valor de epIV
	fmt.Println("Star Wars: Ep.IV - Marco %d", epIV)
}
