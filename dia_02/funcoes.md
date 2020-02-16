# Dia 02 - Cont.

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

[Interfaces :fast_forward:](interfaces.md#interfaces)