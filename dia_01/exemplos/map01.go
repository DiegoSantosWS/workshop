package main

import "fmt"

func main() {
	naves := make(map[string]string)

	naves["YT-1300"] = "Millennium Falcon"
	naves["T-65"] = "X-Wing"
	naves["RZ-1"] = "A-Wing"
	naves["999"] = "Tunder Tanque"

	fmt.Println("Quantidade de naves:", len(naves))
	fmt.Println(naves)
	fmt.Printf("Nave do Han Solo: %s\n", naves["YT-1300"])

	fmt.Println("999 não é uma nave. Removendo...")
	delete(naves, "999")

	fmt.Println("Quantidade de naves atualizada:", len(naves))
	fmt.Println(naves)
}
