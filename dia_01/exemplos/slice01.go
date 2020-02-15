package main

import "fmt"

func main() {
	// declaracao com var
	var s1 []int
	fmt.Println("Slice 1:", s1)
	// declaração curta
	s2 := []int{}
	fmt.Println("Slice 2:", s2)
	// tamanho de um slice
	fmt.Println("Tamanho do slice 1:", len(s1))
	fmt.Println("Tamanho do slice 2:", len(s2))
}
