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

- Desempenho
- Escalabilidade
- Facilidade de manutenção
- Propósito geral

### Hello World

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

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