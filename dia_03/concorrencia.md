# Dia 03

# Concorrência


# Goroutines

Em Go cada atividade que executa de forma concorrênte é chamada de  *`goroutine`*. 

Normalmente desenvolver software com programação concorrente nunca é simples, mas Go tornou isso mais fácil do que em outras linguagens.
A criação de de uma thread **também chamada de goroutine** é praticamente trivial no dia a dia do desenvolvedor.

No exemplo abaixo podemos ver como é feita a chamada de uma goroutine.

```go
package main

func exempleGoroutine(str string) {
    for i := 1; i < 3; i++ {
        fmt.Printf("%s: %d", str, i)
    }
}

func main() {
    exempleGoroutine("direto") // espera que ela retorne
    go exempleGoroutine("com go routine") //cria uma goroutine e não espara que retorne.
}
```

### Canais (channals)

Se goroutines são atividades de um programa concorrênte, canais(channals) são as conexões entre elas. Um canal é uma sistema de comunicação que permite a uma goroutine enviar valores para outra goroutine.

Canal é um condutor de valores de um tipo particular, chamados de *tipo de elemento* do canal.

```go
package main


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

```

```go
package main


func main() {
	sendDataToChannal()
}

func sendDataToChannal() {
	ch := make(chan int)
	ch <- 1 //enviando dados para um canal
	<-ch

	ch <- 2
	fmt.Println(<-ch)
}

```

# Defer


# WaitGroup