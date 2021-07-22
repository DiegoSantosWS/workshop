package main

import "fmt"

var nome, desc string
var diametro int32
var massa float64
var ativo bool
var terreno []string

func main() {
	nome = "Alderaan"
	desc = "Planeta"
	diametro = 12500
	massa = 1.989E+30
	ativo = false
	terreno = []string{
		"Pradarias",
		"Montanhas",
	}
	fmt.Println(nome)
	fmt.Println(desc)
	fmt.Println("Raio (km)", diametro)
	fmt.Println("Massa (kg)", massa)
	fmt.Println("Ativo? ", ativo)
	fmt.Println("Terreno", terreno)
}
