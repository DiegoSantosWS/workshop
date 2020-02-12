package main

import "fmt"

var (
	nome     string  = "Endor"
	desc     string  = "Lua"
	diametro int32   = 4900
	massa    float64 = 1.024e26
	ativo    bool    = true
	terreno          = []string{
		"Florestas",
		"Montanhas",
		"Lagos",
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
