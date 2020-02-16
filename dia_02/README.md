# Dia 02

## Estruturas de controle

Estruturas de controle são ultilizadas para alterar a forma como o nosso código é executado. Podemos, por exemplo, fazer com que uma parte do nosso código seja repitido várias vezes, ou que seja executado caso uma condição seja satisfeita.

### if

O `if` é uma instrução que avalia uma condição booleana. Para entender melhor como ele funciona, vamos analisar o seguinte problema:

> Em uma soma de dois números, onde `a` = 2 e `b` = 4, avalie o valor de `c`. Se `c` for igual a 6, imprima na tela *"Sua soma está correta!"*, caso contrário, imprima *"Sua soma está errada!"*.

```go
package main 

func main() {
    var a, b = 2, 4
    c := (a + b)
    if c == 6 {
        fmt.Println("Sua soma está correta.")
        return
    }
    //uma forma d fazer se não
    fmt.Println("Sua soma está errada.")
}
```

#### Outro Exemplo:

```go
package main 

func main() {
    if 2%2 == 0 {
        fmt.Println("É par.")
    } else {
        fmt.Println("É impar.")
    }

    if num := 2 num < 0 {
        fmt.Println(num, "É negativo.")
    } else if num < 10 {
        fmt.Println(num, "Tem um dígito.")
    } else {
        fmt.Println(num, "Tem vários dígitos.")
    }
}
```

### Switch

A instrução ***`switch`*** é uma maneira mais fácil de evitar longas instruções ***`if-else`***. Com ela é possível realiza ações diferentes com base nos possíveis valores de uma expressão.

#### Exemplo 1

```go
package main 

func main() {
    i := 2
    switch i {
    case 1:
        fmt.Println("Valor de ", i, " por extenso é: um")
    case 2:
        fmt.Println("Valor de ", i, " por extenso é: dois")
    case 3:
        fmt.Println("Valor de ", i, " por extenso é: três")
    }
}
```

O ***`switch`*** pode testar valores de qualquer tipo, além de podermos usar vírgula para separar várias expreções em uma mesma condição ***`case`***.

#### Exemplo 2

```go
package main

func main() {
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("É fim de semana.")
    default:
        fmt.Println("É dia de semana.")
    }
}
```

O ***`switch`*** sem uma expressão é uma maneira alternativa para expressar uma lógica ***`if-else`***.

#### Exemplo 3

```go
package main

func main() {
    j := 3
    switch {
    case 1 == j:
        fmt.Println("Valor por extenso é: um")
    case 2 == j:
        fmt.Println("Valor por extenso é: dois")
    default:
        fmt.Println("Valor não encontrado.")
    }
}
```

### for

Em outras linguagens de programação temos várias formas de fazer laços de repetição, porém, em **Go** só temos uma forma, e é usando a palavra reservada ***`for`***.

#### Exemplo 1

A forma tradicional, que já conhecemos, e que no exemplo vai imprimir números de 1 a 10.

```go
package main

func main(){
    for i := 1; i <=10; i++ {
        fmt.Println("O número é: ", i)
    }
}
```

#### Exemplo 2

```go
package main

func main(){
    i := 5
    for i <= 5 {
        fmt.Println("O número é: ", i)
        i = i + 1
    }
}
```

#### Exemplo 3 loop infinito

```go
package main

func main(){
    for {
        fmt.Println("Olá sou o infinito")
        break
    }
}
```

### for range

Já vimos as outras formas de usar o ***`for`***, agora falta o ***`range`***. Essa expressão espera receber uma lista (`array` ou `slice).

```go

func main() {

    listaDeCompras := []string{"arroz", "feijão", "melancia", "banana", "maçã", "ovo", "cenoura"}
    for k, p := range listaDeCompras {
        retornaNomeVegetal(k, p)
    }
}

func retornaNomeVegetal(key int, str string) {
    switch str {
    case "banana", "cenoura":
       fmt.Println("A sessão", key, " e vegetal é:", str)
    default:
        return
    }
}
```

## Struct

***`Stuct`*** é um tipo de dado agregado que agrupa zero ou mais valores nomeados de tipo quaisquer como uma única entidade. Cada valor é chamado de ***campo***.

### Struct Nomeada

Uma `struct` nomeada recebe um nome em sua declaração. 
Para exemplificar, criaremos uma `strutc` para representar um cadastro de funcionário em uma empresa. Seus campos pode ser acessados través da expressão `variaval.Name`, exemplo:

```go
package main

type Employee struct {
    ID      int
    Name    string
    Age     *time.Time
    Salary  float64
    Company string
}

func main() {
    cl := Employee{}
    //forma de acesso
    cl.ID = 1
    cl.Name = "Diego dos Santos"
    cl.Age = nil
    cl.Salary = 100.55
    cl.Company = "Fliper"
    fmt.Println("o nome é:", cl.Name, " trabalha na empresa: ", cl.Company)
    //oura forma de popular structs
    cl1 := Employee{
        ID:      1,
        Name:    "Francisco Oliveira",
        Age:     nil,
        Salary:  2000.50,
        Company: "Iron Mountain",
    }
    fmt.Println("o nome é:", cl1.Name, " trabalha na empresa: ", cl1.Company)

}
```

### Struct anônima

Uma ***`struct`*** anônima é tipo sem um nome como referência. Sua declaração é semelhante a uma ***declaração rapída de variável***.

> Só devemos usar uma `struct` anônima quando não há necessida de criar um objeto para o dado que será transportado por ela.

```go
package main

func main() {
    inferData("Diego", "Santos")
    inferData("Francisco", "Oliveira")
}

func inferData(fN, lN string) {
    name1 := struct{FirstName, LastName string}{FirstName: fN, LastName: lN}
    fmt.Println("O nome é:", name1.FirstName, name1.LastName)
}
```

[Funções :fast_forward:](funcoes.md#funções)
