package main

import "fmt"

var (
	nome     string  = "Hoth"
	desc     string  = "Planeta"
	diametro int32   = 7200
	massa    float64 = 5.972e+24
	ativo    bool    = true
	terreno          = []string{
		"Tundra",
		"Cavernas de Gelo",
		"Cadeias de Montanhas",
	}
)

func main() {

	fmt.Println(nome)
	fmt.Println(desc)
	fmt.Println("Raio (km)", diametro)
	fmt.Println("Massa (kg)", massa)
	fmt.Println("Ativo? ", ativo)
	fmt.Println("Terreno", terreno)
}
