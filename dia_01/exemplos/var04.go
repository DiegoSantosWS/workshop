package main

import "fmt"

func main() {
	nome := "Endor"
	desc := "Lua"
	diametro := 4900
	massa := 1.024e26
	ativo := true
	terreno := []string{
		"Florestas",
		"Montanhas",
		"Lagos",
	}

	fmt.Println(nome)
	fmt.Println(desc)
	fmt.Println("Raio (km)", diametro)
	fmt.Println("Massa (kg)", massa)
	fmt.Println("Ativo? ", ativo)
	fmt.Println("Satelites", terreno)
}
