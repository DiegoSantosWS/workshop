# Dia 02

## Estrutura de controle

Estrutura de controle é ultilizada para alterar a forma em que nosso código é executado. Podemos fazer, por exemplo, que uma parte do nosso código se repita varias vezes, ou que seja executada um pedaço de código caso uma condição seja apresentada.

### if

O <code>***if***</code> é uma comando que representa uma condição, para entender melhor como ele funciona veja o problema.

> Em uma soma de dois números onde a = 2 e b = 4 descubra o valo de c, se o c for igual a 6 imprima na tela "Sua soma está correta."

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

```go
package main 

func main() {
    var a, b = 2, 4
    c := (a + b)
    if c == 7 {
        fmt.Println("Sua soma está correta.")
    } else {
        //uma forma de fazer o else
        fmt.Println("Sua soma está errada.")
    }
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

A instrução <code>***switch***</code> expressam condições em através de vários ramos. em ***_Go*** podemos usar vírgula para separar várias expreções em uma mesma condição <code>***case***</code>.

O <code>***switch***</code> sem nenhuma expreção é uma logical para representar <code>***if/else***</code>

# Exemplo 1

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

# Exemplo 2

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

# Exemplo 3

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

# for

Em outras linguagens temos várias formas de fazer laços de repetição, mas é _Go so temos uma forma, e é usando a palavra reservada <code>***for***</code>.

### Exemplo 1

Na forma tradicional que ja conhecemos vamos imprimir numeros de 1 a 10.

```go
package main

func main(){
    for i := 1; i <=10; i++ {
        fmt.Println("O número é: ", i)
    }
}
```

# Exemplo 2

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

# Exemplo 3 loop infinito

```go
package main

func main(){
    for {
        fmt.Println("Olá sou o infinito")
        break
    }
}
```

# for range

Já vimos as outras formas de usar o <code>***for***</code>, agora falta o <code>***range***</code>. Essa expressão espera recever uma lista "array ou slice"

```go

func main() {

    listaDeCompras := []string{"arroz", "fejão", "melancia", "banana", "maçã", "ovo", "cenora"}
    for k, p := range listaDeCompras {
        retornaNomeVegetal(k, p)
    }
}

func retornaNomeVegetal(key int, str string) {
    switch str {
    case "banana", "cenora":
       fmt.Println("A sessão", key, " e vejetal é:", str)
    default:
        return
    }
}
```

# Structs

<code>***Stucts***</code> é um tipo de dado agregado que agrupa zero ou mais valores nomeados de tipo quaisquer como uma única entidade. Cada valor é chamado de ***campo***.

### Struct Nomeada

Estrutura no nomeada recebem um nome em sua declaração, para exemplificar criaremos uma strutc para representar um cadastro de funcionário em uma empresa. Seu campos pode ser acessados través da expressão **variaval.Name** veja no exemplo.

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

