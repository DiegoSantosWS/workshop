package main

import (
	"fmt"
)

const (
	estrelaAna byte = iota
	estrelaSubAna
	estrelaAnaBranca
	estrelaAnaVermelha
	estrelaAnaMarrom
)

func main() {
	fmt.Println(estrelaAna)
	fmt.Println(estrelaSubAna)
	fmt.Println(estrelaAnaBranca)
	fmt.Println(estrelaAnaVermelha)
	fmt.Println(estrelaAnaMarrom)
}
