package main

import "fmt"

var nome, desc string = "Tatooine", "Planeta"
var diametro int32 = 10465
var massa float64 = 5.972E+24
var ativo bool = true
var terreno = []string{
    "Deserto",
}

func main() {
	fmt.Println(nome)
	fmt.Println(desc)
	fmt.Println("Diâmetro (km)", diametro)
	fmt.Println("Massa (kg)", massa)
	fmt.Println("Terreno", terreno)
}
