# Dia 01

## O que é Go

> _Go é uma linguagem de programação de código aberto que facilita a criação de software **simples**, **confiável** e **eficiente**._
>> <cite>[golang.org][1]</cite>

### Um pouco de história

O projeto da linguagem Go foi iniciado em 2007 pelos engenheiros da Google Rob Pike, Ken Thompson e Robert Griesemer. A linguagem foi apresentada pela primeira vez em 2009 e teve a versão 1.0 lançada em 2012.

### Semântica

Go tem uma certa semelhança com C, principalmente no que diz respeito a alcançar o máximo de efeito com o mínimo de recursos.

Porém, ela não é uma versão atualizada do C. Na verdade, Go adapta boas ideia de várias linguagens, sempre descartando funcionalidades que trazem complexidade e código não confiável.

É Go compilada e estaticamente tipada.

Possui ponteiros, mas não possui aritmética de ponteiros.
É uma linguagem moderna que utiliza recursos para concorrência novos e eficientes.
Além de ter gerenciamento de memória automático, também conhecido como garbage collection (coleta de lixo)

### Por quê Go?

#### Código aberto
Go é uma linguagem de código aberto, o que significa que não é restritiva e qualquer pessoa pode contribuir.

#### Fácil de aprender
A sintaxe de Go é pequena em comparação com outras linguagens. É muito limpa e fácil de aprender e até mesmo desenvolvedores de outras linguagens, familiarizados com C, podem aprender e entender Go facilmente.

#### Desempenho rápido
A pequena sintaxe e o modelo de concorrência do Go o tornam uma linguagem de programação muito rápida. Go é compilada em código de máquina e seu processo de compilação também é muito rápido.

#### Modelo de Concorrência Fácil
Go é construído para concorrência, o que facilita a execução de várias tarefas ao mesmo tempo. Go possui _**goroutines**_, threads leves que se comunicam através de um canais.

#### Portabilidade e multiplataforma
Go é uma linguagem multiplataforma. Pode-se escrever código facilmente em qualquer ambiente (OSX, Linux ou Windows). Portanto, o código escrito no Windows pode ser compilado e distribuído em um ambiente Linux.

#### Design explícito para a nuvem
Go foi escrito especialmente para a nuvem. Em Go todas as bibliotecas e dependências são vinculadas em um único arquivo binário, eliminando assim a instalação dependências nos servidores. E esse é outro motivo para o seu crescimento e popularidade.

#### Segurança
Como o Go é estaticamente e fortemente tipada, isso implica que você precisa ser explícito no tipo de dados que está passando e também significa que o compilador conhece o tipo de cada variável, respectivamente.

#### Coleta de lixo
Go possui um coletor de lixo (_Garbage Colletion_) para gerenciamento automático de memória. Esse recurso de GC faz a alocação e a remoção de objetos sem nenhuma pausa e, portanto, aumenta a eficiência das aplicações.

#### Biblioteca Padrão Poderosa
A biblioteca padrão muito é poderosa e cheia de recursos. Com ela é possível facilmente construir um servidor Web, manipular E/S, criptografia e assim por diante.


### Hello World

```go
// Ex1
package main

import "fmt"

func main() {
    fmt.Println("Hello, World")
}
```

Como podemos testar esse código?

Primeiramente, vamos separar o "domínio" (regras de negócio) do restante do código (efeitos colaterais). A função `fmt.Println` é um efeito colateral (que está imprimindo um valor no `stdout` - saída padrão) e a `string` que estamos passando para ela é o nosso domínio.

```go
// Ex2
package main

import "fmt"

func Hello() string {
    return "Hello, World"
}

func main() {
    fmt.Println(Hello())
}
```

Criamos uma nova função: `Hello`, mas dessa vez adicionamos a palavra `string` na sua definição. Isso significa que essa função retornará uma `string`.

```go
// Ex3
package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello()
    want := "Hello, World"

    if got != want {
        t.Errorf("Got '%s', want '%s'", got, want)
    }
}
```

Esse é o código que criaremos para testar a nossa função `Hello()`.
Vamos criar um arquivo chamado `hello_test.go` e nele coloque este código.

Percebam que não é preciso usar vários _frameworks_ (ou bibliotecas) de testes. Tudo o que precisamos está pronto na linguagem e a sintaxe é a mesma para o resto dos códigos que você irá escrever.

Escrever um teste é como escrever uma função, com algumas regras:

- O código precisa estar em um arquivo que termine com **_test.go**
- A função de teste precisa começar com a palavra **Test**
- A função de teste recebe um único argumento, que é `t *testing.T`

Por enquanto, isso é o bastante para saber que o nosso `t` do tipo `*testing.T` é a nossa porta de entrada para a ferramenta de testes.

## Tipos básicos

### String
### Números
### boleanos

## Variáveis / Constantes / Ponteiros

## Tipos Compostos

### Array
### Slices
### Maps


[1]:https://golang.org/