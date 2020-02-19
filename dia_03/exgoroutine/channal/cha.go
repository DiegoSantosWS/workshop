package main

import "fmt"

func main() {
	sendDataToChannal()
}

func sendDataToChannal() {
	ch := make(chan int, 1)
	ch <- 1 //enviando dados para um canal
	<-ch

	ch <- 2
	fmt.Println(<-ch)
}
