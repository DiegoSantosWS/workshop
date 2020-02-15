package main

import "fmt"

func main() {
	// Naves do jogo "Star Wars: Battlefront"
	rebeldes := [...]string{"'X-Wing'", "'A-Wing'", "'Millenium Falcon'"}
	imperiais := [...]string{"'TIE Fighter'", "'TIE Interceptor'", "'Imperial Shuttle'", "'Slave I'"}

	naves := make([]string, 0, 0)
	fmt.Printf("Cap: %d - %v\n", cap(naves), naves)
	naves = append(naves, "''")
	fmt.Printf("Cap: %d - %v\n", cap(naves), naves)
	naves = append(naves, rebeldes[:]...)
	fmt.Printf("Cap: %d - %v\n", cap(naves), naves)
	naves = append(naves, imperiais[:]...)
	fmt.Printf("Cap: %d - %v\n", cap(naves), naves)
}
