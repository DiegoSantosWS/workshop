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

Já vimos as outras formas de usar o ***`for`***, agora falta o ***`range`***. Essa expressão espera receber uma lista (`array` ou `slice`).

```go
func exemploFor4() {
	listaDeCompras := []string{"arroz", "feijão", "melancia", "banana", "maçã", "ovo", "cenoura"}
	for k, p := range listaDeCompras {
		retornaNomeFruta(k, p)
	}
}

func retornaNomeFruta(key int, str string) {
	switch str {
	case "melancia", "banana", "maçã":
		fmt.Println("Na posição", key, "temos a fruta:", str)
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

## Funções

Funções são pequenas unidades de códigos que podem abistrair ações, retornar e/ou receber valores.

### Como declarar uma função?

Declarar uma função é algo bem simples, utilizando a palavra reservada ***`func`*** seguida do identificador.

```go
package main

func nomeDaFuncao(){}

```

Essa é a declaração mais simples de função que temos. No exemplo acima criamos uma função que não recebe nenhum parametro e não retorna nada, o nome dela poderia ser `fazNada`.

Uma função em Go também é um tipo e pode ser declarada da segunte forma:

```go
package main

type myFunc = func(l, b int) int

func main() {
	soma(func(l, b int) int {
		return l + b
	})
}

func soma(fn myFunc) {
	res := fn(1, 3)
	fmt.Println(res)
}

```

### Declaração de função que recebe parâmetros

Podemos declarar uma função que recebe dois números e faz uma multiplicação.

```go
package main

func main() {
	fmt.Println("Resultado é:", multiplica(3, 8))
}

func multiplica(a, b int) int {
	return (a * b)
}
```

Veja que na declaração da `func multiplica(a, b int)` os parametros foram passados um seguido do outro, isso por que eles são do mesmo **tipo** (`int`). Caso fossem de tipos diferentes seria nessário declarar cada tipo separadamente, exemplo `func minhaFunc(str string, i int)`.

### Funções anônimas

Go também suporta declaração de funções anônimas. Funções anônimas são úteis quando você deseja definir uma função em linha sem ter que nomeá-la.

```go
package main

func main() {
    fn := exemploAnonimo()
    fmt.Println("Resultado é:", fn)
    fmt.Println("Resultado é:", fn)
    fmt.Println("Resultado é:", fn)
}

func exemploAnonimo() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
```

### Função com retorno nomeado

Podemos criar uma função e nomear o retorno da mesma. veja o exemplo:

```go
package main

func main() {
    fn := exemploNomeado()
    fmt.Println("Nome é:", exemploNomeado("Marcela"))
    fmt.Println("Nome é:", exemploNomeado("Diego"))
    fmt.Println("Nome é:", exemploNomeado("Francisco"))
}

func exemploNomeado(str string) (nome string) {
    nome = str
    return 
}
```

### Funções variádicas

Função variádica é uma função que pode receber qualquer número de argumentos à direita e de um mesmo tipo. Um bom exemplo de função variádica é a função `fmt.Println`. A função pode ser chamada de forma usual, com argumentos individuais ou uma lista (`Slice`).

```go
package main

func main() {
    fmt.Println("Resultado é:", exemploVariadico(1,2))
    fmt.Println("Resultado é:", exemploVariadico(2,3))
    fmt.Println("Resultado é:", exemploVariadico(3,4))
}

func exemploVariadico(numeros ...int) (total int) {
    total = 0

    for _, n := range numeros {
        total += n
    }
	return 
}
```

## Erros

**Erros** são um assunto muito complexo em Go, pois não existe um tratamento de exeção como em outras linguagens. A única forma de se tratar *erros* em Go é usando a condição *if* ou então podemos criar uma função para realizar o tratamento. veja os exemplos:


```go
package main

func main() {
    tot, err := exemploVariadicoWithErr(1,2)
    if err != nil {
        return
    }
    fmt.Println("Resultado é:", tot)

    tot2, err := exemploVariadicoWithErr(2,3)
    if err != nil {
        return
    }
    fmt.Println("Resultado é:", tot2)

    tot3, err := exemploVariadicoWithErr(3,4)
    checkErr(err)
    fmt.Println("Resultado é:", tot3)
}

func exemploVariadicoWithErr(numeros ...int) (total int, err error) {
    total = 0

    for _, n := numeros {
        total += n
    }
    if total == 0 {
        err = errors.New("O resultado não pode ser zero")
        return
    }
	return 
}

func checkErr(err error) {
    if err != nil {
        return
    }
}
```
Como mostrado no exemplo acima, uma função pode retornar algum resultado e/ou erro.

## Métodos

Métodos em Go são uma variação da declaração de função. No método, um parâmetro extra aparece antes do nome da função e é chamado de receptor (*receiver*).

Métodos podem ser definidos para qualquer tipo de receptor, até mesmo ponteiros, exemplo:


```go
package main

type area struct {
    Largura int
    Altura  int
}

func (r *area) CalculaArea() int {
    res := r.Largura * r.Altura
    return res
}

func (r area) CalculaPerimetro() int {
    res := 2*r.Largura * 2*r.Altura
    return res
}


func main() {
    a := area{Largura: 10, Altura: 5}
    resultArea := a.CalculaArea()
    fmt.Println("area: ", resultArea)
    perim := &a //repassando os valores
    resultPerim := perim.CalculaPerimetro()
    fmt.Println("perim: ", resultPerim)
}
```

## Interfaces

Uma `interface` é uma coleção de metódos que um tipo concreto deve implementar para ser considerado uma instância dessa interface. Portanto, uma `interface` define, mas, não declara o comportamento do tipo.

Para exemplificar vamos usar os mesmos exemplos que usamos para criar métodos.

```go
package main

import "fmt"

// Geo interface base para figuras geométricas
type Geo interface {
	Area() float64
}

// Retangulo representa um retângulo
type Retangulo struct {
	Largura float64
	Altura  float64
}

// Area calcula a are de um retângulo
func (r *Retangulo) Area() float64 {
	res := r.Largura * r.Altura
	return res
}

// Triangulo representa um triângulo
type Triangulo struct {
	Base   float64
	Altura float64
}

// Area calcula a are de um triângulo
func (t *Triangulo) Area() float64 {
	res := (t.Base * t.Altura) / 2
	return res
}

func imprimeArea(g Geo) {
	fmt.Println(g)
	fmt.Println(fmt.Sprintf("Área      : %0.2f", g.Area()))
}

func main() {
	r := Retangulo{
		Altura:  10,
		Largura: 5,
	}

	t := Triangulo{
		Base:   10,
		Altura: 5,
	}

	imprimeArea(&r)
	imprimeArea(&t)
}
```
