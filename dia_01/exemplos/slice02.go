package main

import "fmt"

func main() {
	// Naves do jogo "Star Wars: Battlefront"
	naves := [...]string{
		1: "X-Wing",
		2: "A-Wing",
		3: "Millenium Falcon",
		4: "TIE Fighter",
		5: "TIE Interceptor",
		6: "Imperial Shuttle",
		7: "Slave I",
	}
	// cria um slice de naves[1] até naves[3]
	rebeldes := naves[1:4]
	fmt.Println(rebeldes)
}
