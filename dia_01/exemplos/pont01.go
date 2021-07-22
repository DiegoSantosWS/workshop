package main

import "fmt"

func main() {
	var p *int
	eraOuroSith, epIV := 42, 37
	// ponteiro para eraOuroSith
	p = &eraOuroSith
	// valor de eraOuroSith por meio do ponteiro
	fmt.Printf("Era de Ouro dos Sith - %d anos antes do Ep.IV (%#x)\n", *p, p)
	// atualiza o valor de eraOuroSith por meio do ponteiro
	*p = 5000
	// o novo valor de eraOuroSith
	fmt.Printf("Era de Ouro dos Sith - %d anos antes do Ep.IV | Atualizado (%#x)\n", *p, p)
	// ponteiro para epIV
	p = &epIV
	// divide epIV por meio do ponteiro
	*p = *p / 38
	// o novo valor de epIV
	fmt.Printf("Star Wars: Ep.IV é o Marco %d (%#x)\n", epIV, p)
}
