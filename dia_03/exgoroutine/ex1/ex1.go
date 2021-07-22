package main

import (
	"fmt"
	"time"
)

func exempleGoroutine(str string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%s: %d\n", str, i)
	}
}

func main() {
	exempleGoroutine("direto")
	go exempleGoroutine("com go routine")
	time.Sleep(time.Second * 1)
}
