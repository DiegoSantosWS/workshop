# Funções

> Funções são pequenas unidades de códigos que podem abistrair ações, retornar e/ou receber valores.

### Como declarar uma função?

Declarar uma funão é algo bem simples, ultilizando a palavra reservada **func** seguida do identificador.

```go
package main

func nomeDaFuncao(){}

```

Essa é a declaração mais simples de função que temos. No exemplo acima criamos uma função que não recebe nenhum parametro e não retorna nada, o nome dela poderia ser `faz_nada`.

Uma função em Go também é um tipo podendo ser declarada da segunte forma.

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

Veja que na declaração da **func multiplica** eu passei os parametros um seguido do outro, isso por que eles são do mesmo *type (tipo)*, se fossem de tipos diferentes seria nessário declarar cada parametro separadamente. **func multiplica(str string, i int)**.

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

Podemos criar uma função e nomear o retorno da mesma. veja o exemplo

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

Função variádica pode receber qualque numero de argumentos à direita de um mesmo tipo, um bom exemplo de função variádica é a func fmt.Println. Função pode ser chamada de forma usual, com argumentos individuais, caso queira passar uma lista *Slice* pode ser aplicado como mostra no exemplo à seguir.

```go
package main

func main() {
    fmt.Println("Resultado é:", exemploVariadico(1,2))
    fmt.Println("Resultado é:", exemploVariadico(2,3))
    fmt.Println("Resultado é:", exemploVariadico(3,4))
}

func exemploVariadico(numeros ...int) (total int) {
    total = 0

    for _, n := numeros {
        total += n
    }
	return 
}
```

# Erros

Erros é um assunto muito complexo em Go pois não existe um tratamento de exeção como em outras linguagens, a unica forma de se tratar *error* em Go é usando a condição *if* para isso ou então podemos criar uma função para realizar o tratamento. veja os exemplos.


```go
package main

func main() {
    tot, err := exemploVariadico(1,2)
    if err != nil {
        return
    }
    fmt.Println("Resultado é:", tot)

    tot2, err := exemploVariadico(2,3)
    if err != nil {
        return
    }
    fmt.Println("Resultado é:", tot2)

    tot3, err := exemploVariadico(3,4)
    checkErr(err)
    fmt.Println("Resultado é:", tot3)
}

func exemploVariadico(numeros ...int) (total int, err error) {
    total = 0

    for _, n := numeros {
        total += n
    }
    if total == 0 {
        err = errors.New("O resoltado não pode ser zero")
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
Como mostrado no exemple acima uma função pode returnar algum resultado e/ou errors.

# Metódos

Metódos em Go são definidos por um tipo struct, metódos pode ser definidos para qualquer tipo de receptor seja ponteiro ou valor. veja exemplo


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