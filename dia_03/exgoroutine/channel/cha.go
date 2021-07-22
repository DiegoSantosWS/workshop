package main

import (
	"fmt"
	"time"
)

func main() {
	sendDataToChannel()

	c := make(chan int)
	go doisTresQuatroVezes(2, c)

	a, b := <-c, <-c // recebendo os dados do canal
	fmt.Println(a, b)

	fmt.Println(<-c)
}

func sendDataToChannel() {
	ch := make(chan int, 1)
	ch <- 1 //enviando dados para um canal
	<-ch

	ch <- 2
	fmt.Println(<-ch)
}

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}
