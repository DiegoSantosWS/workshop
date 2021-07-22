package main

import "fmt"

var nome, desc = "Yavin IV", "Lua"
var diametro = 10200
var massa = 641693000000000.0
var ativo = true
var terreno = []string{
	"Selva",
	"Florestas Tropicais",
}

func main() {
	fmt.Println(nome)
	fmt.Println(desc)
	fmt.Println("Raio (km)", diametro)
	fmt.Println("Massa (kg)", massa)
	fmt.Println("Terreno", terreno)
}
