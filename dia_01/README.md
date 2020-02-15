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
Go foi escrito especialmente para a nuvem. Em Go todas as bibliotecas e dependências são vinculadas em um único arquivo binário, eliminando assim a instalação de dependências nos servidores. E esse é outro motivo para o seu crescimento e popularidade.

#### Segurança
Como o Go é estaticamente e fortemente tipada, isso implica que você precisa ser explícito no tipo de dados que está passando e também significa que o compilador conhece o tipo de cada variável, respectivamente.

#### Coleta de lixo
Go possui um coletor de lixo (_Garbage Colletion_) para gerenciamento automático de memória. Esse recurso de GC faz a alocação e a remoção de objetos sem nenhuma pausa e, portanto, aumenta a eficiência das aplicações.

#### Biblioteca Padrão Poderosa
A biblioteca padrão muito é poderosa e cheia de recursos. Com ela é possível facilmente construir um servidor Web, manipular E/S, criptografia e é claro, criar testes, dos quais falaremos mais ao longo do workshop.


### Hello World

```go
// Ex1
package main

import "fmt"

func main() {
    fmt.Println("Hello, World")
}
```

Assim como em outras linguagens... em Go temos também o clássico **Hello World**.

Mas... como podemos testar esse código?

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

- O código precisa estar em um arquivo que termine com **_test.go**.
- A função de teste precisa começar com a palavra **Test**.
- A função de teste recebe um único argumento, que é `t *testing.T`.

Por enquanto, isso é o bastante para saber que o nosso `t` do tipo `*testing.T` é a nossa porta de entrada para a ferramenta de testes.

## Tipos básicos

Go apresenta várias maneiras de organizar dados, desde tipos que correspondem aos recursos do hardware até tipo convenientes para a representação de estruturas de dados complexas.

### String

Uma `string` é uma sequência imutável de `bytes`. Podem conter qualquer dado, mas normalmente contêm texto legível aos seres humanos.

Strings são convencionalmente interpretadas como sequências de pontos de código **Unicode** (`runas`) codificados em **UTF-8**.

### Números

Os tipos numéricos em Go incluem vários tamanhos de inteiros, ponto flutuante e números complexos.

#### Inteiros

- int8, uint8, int16, uint16, int32, uint32, int64, uint64
- int, uint (assume o tamanho especificado pelo compilador)
- byte: sinônimo para uint8
- runa: sinônimo para int32
- uintptr: tipo sem um tamanho espeficidado (usado em programação de baixo nível)

#### Ponto Flutuante

- float32, float64
- Segue o padrão IEEE 754

#### Complexos

- complex64, complex128 - Podem ser criados pela função `complex`

### Booleanos

- true, false

## Variáveis / Constantes / Ponteiros

### Variáveis

Go é uma linguagem fortemente tipada, o que implica que todas as variáveis são elementos nomeados que estão vinculados a um valor e um tipo.

A forma longa para declarar uma variável em Go segue o seguinte formato:

```go
var <lista de identificadores> <tipo>
```

A palavra-chave `var` é usada para declarar um ou mais identificadores de variáveis, seguidos do seus respectivos tipos. O trecho de código a seguir mostra a declaração de diversas variáveis:

```go
// var01.go
...
var nome, desc string
var diametro int32
var massa float64
var ativo bool
var terreno []string
...
```

#### O valor zero

O trecho de código anterior mostra vários exemplos de variáveis sendo declaradas com uma variedade de tipos. À primeira vista, parece que essas variáveis não têm um valor atribuído. Na verdade, isso contradiz nossa afirmação anterior de que todas as variáveis em Go estão vinculadas a um tipo e um valor.

Durante a declaração de uma variável, se um valor não for fornecido, o compilador do Go vinculará automaticamente um valor padrão (ou um valor zero) à variável para a inicialização adequada da memória.

A tabela a seguir mostra os tipos do Go e seus valores zero padrão:

Tipo | Valor zero
-----|-----
string| "" (string vazia)
Numérico – Inteiro: byte, int, int8, int16, int32, int64, rune, uint, uint8, uint16, uint32, uint64, uintptr| 0
Numérico – Ponto flutuante: float32, float64| 0.0
booleano| false
Array| Cada índice terá um valor zero correspondente ao tipo do array.
Struct| Em uma estrutura vazia, cada membro terá seu respectivo valor zero.
Outros tipos: Interface, função, canais, slice, mapas e ponteiros| nil

#### Declaração inicializada

Go também suporta a combinação de declaração de variável e inicialização como uma expressão usando o seguinte formato:

```go
var <lista de identificadores> <tipo> = <lista de valores ou expressões de inicialização>
```

O seguinte trecho de código mostra a combinação de declaração e inicialização:

```go
// var02.go
...
var nome, desc string = "Tatooine", "Planeta"
var diametro int32 = 10465
var massa float64 = 5.972E+24
var ativo bool = true
var terreno = []string{
    "Deserto",
}
...
```

#### Omitindo o tipo das variáveis

Em Go, também é possível omitir o tipo, conforme mostrado no seguinte formato de declaração:

```go
var <lista de identificadores> = <lista de valores ou expressões de inicialização>
```

O compilador do Go irá inferir o tipo da variável com base no valor ou na expressão de inicialização do lado direito do sinal de igual, conforme mostrado no trecho de código a seguir.

```go
// var03.go
...
var nome, desc = "Yavin IV", "Lua"
var diametro = 10200
var massa = 641693000000000.0
var ativo = true
var terreno = []string{
    "Selva",
    "Florestas Tropicais",
}
...
```

Quando o tipo da variável é omitido, as informações de tipo são deduzidas do valor atribuído ou do valor retornado de uma expressão.

A tabela a seguir mostra o tipo que é inferido dado um valor literal:

Valor Literal | Tipo inferido
--------------|--------------
Texto com aspas duplas ou simples:<br>`"Lua Yavin IV"`<br><code>\`Sua superfície tem<br>seis continentes ocupando<br>67% do total.\`</code> | string
Inteiros:<br>`-51`<br>`0`<br>`1234` | int
Decimal:<br>`-0.12`<br>`1.0`<br>`1.3e5`<br>`5e-11` | float64
Números complexos:<br>`-1.0i`<br>`2i`<br>`(0+2i)` | complex128
Booleanos:<br>`true`<br>`false` | bool
Arrays:<br>`[2]int{-3, 51}`| O tipo do `array` definido pelo valor literal. Neste caso `[2]int`
Map:<br>`map[string]int{`<br>`"Tatooine": 10465,`<br>`"Alderaan": 12500,`<br>`"Yavin IV": 10200,`<br>`}` | O tipo do `map` definido pelo valor literal. Neste caso `map[string]int`
Slice:<br>`[]int{-3, 51, 134, 0}` | O tipo do `slice` definido pelo valor literal: `[]int`
Struct:<br>`struct{`<br>`nome string`<br>`diametro int`<br>`}{`<br>`"Tatooine", 10465,`<br>`}` | O tipo do `struct` definido conforme o valor literal. Neste caso: `struct{nome string; diametro int}`
Function:<br>`var sqr = func (v int) int {`<br>`    return v * v`<br>`}` | O tipo de `function` definido na definição literal da função. Neste caso, a variável `sqr` terá o tipo: `func (v int) int`

#### Declaração curta de variável

Em Go é possível reduzir ainda mais a sintaxe da declaração de variáveis. Neste caso, usando o formato _short variable declaration_. Nesse formato, a declaração perde a palavra-chave `var` e a especificação de tipo e passa a usar o operador `:=` (dois-pontos-igual), conforme mostrado a seguir:

```go
<lista de identificadores> := <lista de valores ou expressões de inicialização>
```

O techo de código a seguir mostra como usá-la:

```go
// var04.go
...
func main() {
	nome := "Endor"
	desc := "Lua"
	diametro := 4900
	massa := 1.024e26
	ativo := true
	terreno := []string{
        "Florestas", 
        "Montanhas", 
        "Lagos",
    }
...
```

##### Restrições na declaração curta de variáveis

Existem algumas restrições quando usamos a declaração curta de variáveis e é muito importante estar ciente para evitar confusão:

- Em primeiro lugar, ela só pode ser usada dentro de um bloco de funções;
- o operador `:=` declara a variável e atribui os valores;
- `:=` não pode ser usado para atualizar uma variável declarada anteriormente;

#### Declaração de variável em bloco

A sintaxe do Go permite que a declaração de variáveis seja agrupada em blocos para maior legibilidade e organização do código. O trecho de código a seguir mostra a reescrita de um dos exemplos anteriores usando a declaração de variável em bloco:

```go
// var05.go
var (
	nome     string  = "Endor"
	desc     string  = "Lua"
	diametro int32   = 4900
	massa    float64 = 1.024e26
	ativo    bool    = true
	terreno          = []string{
		"Florestas",
		"Montanhas",
		"Lagos",
	}
)
```

### Constantes

Uma constante é um valor com uma representação literal de uma string, um caractere, um booleano ou números. O valor para uma constante é estático e não pode ser alterado após a atribuição inicial.

#### Constantes tipadas

Usamos a palavra chave `const` para indicar a declaração de uma constante. Diferente da declaração de uma variável, a declaração deve sempre incluir o valor literal a ser vinculado ao identificador, conforme mostrado a seguir:

```go
const <lista de identificadores> tipo = <lista de valores ou expressões de inicialização>
```

O seguinte trecho de código mostra algumas constantes tipadas sendo declaradas:

```go
// const01.go
...
const a1, a2 string = "Workshop", "Go"
const b rune = 'G'
const c bool = false
const d int32 = 2020
const e float32 = 2.020
const f float64 = math.Pi * 2.0e+3
const g complex64 = 20.0i
const h time.Duration = 20 * time.Second
...
```

Note que cada constante declarada recebe explicitamente um tipo. Isso implica que a constantes só podem ser usada em contextos compatíveis com seus tipos. No entanto, isso funciona de maneira diferente quando o tipo é omitido.

#### Constantes não tipadas

Constantes são ainda mais interessantes quando não são tipada. Uma constante sem tipo é declarada da seguinte maneira:

```go
const <lista de identificadores> = <lista de valores ou expressões de inicialização>
```

Neste formato, a especificação de tipo é omitida na declaração. Logo, uma constante é meramente um bloco de bytes na memória sem qualquer tipo de restrição de precisão imposta. A seguir, algumas declarações de constantes não tipificadas:

```go
// const02.go
...
const i = "G é" + " para Go"
const j = 'G'
const k1, k2 = true, !k1
const l = 111*100000 + 20
const m1 = math.Pi / 3.141592
const m2 = 1.41421356237309504880168872420969807856967187537698078569671875376
const m3 = m2 * m2
const m4 = m3 * 20.0e+400
const n = -5.0i * 20
const o = time.Millisecond * 20
...
```

A constante `m4` recebe um valor muito grande (`m3 * 20.0e+400`) que é armazenado na memória sem qualquer perda de precisão. Isso pode ser útil em aplicações onde realizar cálculos com um alto nível de precisão é extremamente importante.

#### Atribuindo constantes não tipadas

Mesmo Go sendo uma linguagem fortemente tipada, é possível atribuir uma constante não tipada a diferentes tipos de precisão diferentes, embora compatíveis, sem qualquer reclamação do compilador, conforme mostrada a seguir:

```go
// const03.go
...
const m2 = 1.41421356237309504880168872420969807856967187537698078569671875376
var u1 float32 = m2
var u2 float64 = m2
u3 := m2
...
```

O exemplo anterior mostra a constante não tipada `m2` sendo atribuída a duas variáveis de ponto flutuante com diferentes precisões, `u1` e `u2`, e a uma variável sem tipo, `u3`. Isso é possível porque a constante `m2` é armazenada como um valor não tipado e, portanto, pode ser atribuída a qualquer variável compatível com sua representação (um ponto flutuante).

Como `u3` não tem um tipo específico, ele será inferido a partir do valor da constante, e como `m2` representa um valor decimal, o compilador irá inferir seu tipo padrão, um `float64`.

A declaração de constantes também podem ser organizadas em blocos, aumentando a legibilidade do código, conforme a seguir:

```go
// const04.go
...
const (
	a1, a2 string        = "Workshop", "Go"
	b      rune          = 'G'
	c      bool          = false
	d      int32         = 2020
	e      float32       = 2.020
	f      float64       = math.Pi * 20.0e+3
	g      complex64     = 20.0i
	h      time.Duration = 20 * time.Second
)
...
```

### Enumerações

Um interessante uso para constantes é na criação de enumerações. Usando a declaração de blocos, é facilmente possível criar valore inteiros que aumentam numericamente. Para isso, basta atribuir o valor constante pré-declarado `iota` a um identificador de constante na declaração de bloco, conforme mostrado no exemplo a seguir:

```go
// enum01.go
...
const (
	estrelaHiperGigante = iota
	estrelaSuperGigante
	estrelaBrilhanteGigante
	estrelaGigante
	estrelaSubGigante
	estrelaAna
	estrelaSubAna
	estrelaAnaBranca
	estrelaAnaVermelha
	estrelaAnaMarrom
)
...
```

Nessa situação, o compilador fará o seguinte:

- Declarar cada membro no bloco como um valor constante inteiro não tipado;
- Inicializar a `iota` com o valor zero;
- Atribuir a `iota`, ou zero, ao primeiro membro (`EstrelaHiperGigante`);
- Cada constante subsequente recebe um `int` aumentado em um.

Assim, as constantes da lista receberão os valores de zero até nove.

É importante ressaltar que, sempre que `const` aparecer em um bloco de declaração, o contador é redefinido para zero. No trecho de código seguinte, cada conjunto de constantes é enumerado de zero a quatro:

```go
// enum02.go
...
const (
	estrelaHiperGigante = iota
	estrelaSuperGigante
	estrelaBrilhanteGigante
	estrelaGigante
	estrelaSubGigante
)
const (
	estrelaAna = iota
	estrelaSubAna
	estrelaAnaBranca
	estrelaAnaVermelha
	estrelaAnaMarrom
)
...
```

#### Substituindo o tipo padrão de uma enumeração

Por padrão, uma constante enumerada é declarada como um tipo inteiro não tipado. Porém, podemos substituir o tipo padrão provendo explicitamente um tipo numérico, como mostrado a seguir:

```go
// enum03.go
...
const (
	estrelaAna byte = iota
	estrelaSubAna
	estrelaAnaBranca
	estrelaAnaVermelha
	estrelaAnaMarrom
)
...
```
É possível especificar qualquer tipo numérico que pode representar um inteiro ou um ponto flutuante. No exemplo anterior, cada constante será declarada como um tipo `byte`.

#### Usando `iota` em expressões

Quando a `iota` aparece em uma expressão, o compilador irá aplicar a expressão para cada valor sucessivo. O exemplo a seguir atribui números pares aos membros do bloco de declaração:

```go
// enum04.go
...
const (
	estrelaHiperGigante = 2.0 * iota
	estrelaSuperGigante
	estrelaBrilhanteGigante
	estrelaGigante
	estrelaSubGigante
)
...
```

#### Ignorando valores em enumerações

É possível ignorar certos valores em uma enumeração simplesmente atribuindo a `iota` a um identificador em branco (`_`). No trecho de código a seguir, o valor `0` é ignorado:

```go
// enum05.go
...
const (
	_                   = iota
	estrelaHiperGigante = 1 << iota
	estrelaSuperGigante
	estrelaBrilhanteGigante
	estrelaGigante
	estrelaSubGigante
)
...
```

### Ponteiros

## Tipos Compostos

### Array
### Slices
### Maps


[1]:https://golang.org/