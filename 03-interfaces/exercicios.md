# Etapa 03: Interfaces

Interfaces são o coração do polimorfismo em Go. A grande sacada: você
**não declara** que um tipo implementa uma interface — se ele tem os métodos,
ele já implementa. É o "duck typing" estático.

---

## Exercício 1: Interface Forma

Crie uma interface `Forma` com o método `Area() float64`. Crie um `Retangulo`
que a implemente.

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import "fmt"

type Forma interface {
    Area() float64
}

type Retangulo struct {
    Largura, Altura float64
}

func (r Retangulo) Area() float64 {
    return r.Largura * r.Altura
}

func main() {
    var f Forma = Retangulo{Largura: 3, Altura: 4}
    fmt.Println(f.Area()) // 12
}
```

**Por quê:** `Retangulo` nunca diz `implements Forma`. Como ele tem o método
`Area() float64`, ele **automaticamente** satisfaz a interface. Isso é
implementação implícita — diferente de Java/C# onde você declara explicitamente.
</details>

---

## Exercício 2: Polimorfismo

Adicione um `Circulo` que também implemente `Forma`. Crie uma função
`imprimirArea(f Forma)` que funcione para qualquer forma.

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import (
    "fmt"
    "math"
)

type Forma interface {
    Area() float64
}

type Retangulo struct{ Largura, Altura float64 }
type Circulo struct{ Raio float64 }

func (r Retangulo) Area() float64 { return r.Largura * r.Altura }
func (c Circulo) Area() float64   { return math.Pi * c.Raio * c.Raio }

func imprimirArea(f Forma) {
    fmt.Printf("Área: %.2f\n", f.Area())
}

func main() {
    imprimirArea(Retangulo{Largura: 3, Altura: 4}) // 12.00
    imprimirArea(Circulo{Raio: 5})                 // 78.54
}
```

**Por quê:** `imprimirArea` aceita qualquer coisa que seja `Forma`. Não importa
o tipo concreto — só importa que tenha `Area()`. Esse é o polimorfismo de Go.
</details>

---

## Exercício 3: A interface Stringer

Faça `Circulo` implementar `fmt.Stringer` (método `String() string`) para que
`fmt.Println` o imprima de forma customizada.

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import "fmt"

type Circulo struct{ Raio float64 }

// Stringer é uma interface da stdlib: String() string
func (c Circulo) String() string {
    return fmt.Sprintf("Círculo(raio=%.1f)", c.Raio)
}

func main() {
    c := Circulo{Raio: 2.5}
    fmt.Println(c) // Círculo(raio=2.5)
}
```

**Por quê:** `fmt.Stringer` é uma interface que a própria stdlib procura. Se seu
tipo tem `String()`, o `fmt` usa ela automaticamente ao imprimir. É um exemplo
lindo de como interfaces padronizam comportamento sem acoplamento.
</details>

---

## Exercício 4: Interface vazia e type assertion

Escreva uma função `descrever(i interface{})` que diz se o valor recebido é
um `int`, `string` ou outro tipo.

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import "fmt"

func descrever(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("É um int: %d\n", v)
    case string:
        fmt.Printf("É uma string: %q\n", v)
    default:
        fmt.Printf("Tipo desconhecido: %T\n", v)
    }
}

func main() {
    descrever(42)        // É um int: 42
    descrever("olá")     // É uma string: "olá"
    descrever(3.14)      // Tipo desconhecido: float64
}
```

**Por quê:** `interface{}` (ou `any`, a partir do Go 1.18) aceita qualquer tipo.
O `switch v := i.(type)` é o "type switch": descobre o tipo concreto em tempo
de execução. Use com parcimônia — quase sempre uma interface específica é melhor.
</details>

---

## Exercício 5: Interface de ordenação

Crie um slice de `Retangulo` e ordene-o por área usando `sort.Slice`.

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import (
    "fmt"
    "sort"
)

type Retangulo struct{ Largura, Altura float64 }

func (r Retangulo) Area() float64 { return r.Largura * r.Altura }

func main() {
    formas := []Retangulo{
        {Largura: 3, Altura: 4}, // 12
        {Largura: 1, Altura: 2}, // 2
        {Largura: 5, Altura: 5}, // 25
    }

    sort.Slice(formas, func(i, j int) bool {
        return formas[i].Area() < formas[j].Area()
    })

    for _, f := range formas {
        fmt.Printf("%.0f ", f.Area()) // 2 12 25
    }
    fmt.Println()
}
```

**Por quê:** `sort.Slice` recebe o slice e uma função que define "i vem antes
de j?". Você passa a lógica de comparação como uma função anônima (closure).
É um exemplo de função como argumento — muito comum em Go.
</details>

---

⬅️ Anterior: [Etapa 02](../02-structs-metodos/exercicios.md) · ➡️ Próxima: [Etapa 04 — Concorrência](../04-concorrencia/exercicios.md)
