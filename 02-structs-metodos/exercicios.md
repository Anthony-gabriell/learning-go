# Etapa 02: Structs e Métodos

Como Go organiza dados (structs) e comportamento (métodos). Aqui entra
também a diferença crucial entre receiver por valor e por ponteiro.

---

## Exercício 1: Struct Pessoa

Crie uma struct `Pessoa` com `Nome` (string) e `Idade` (int). No `main`,
crie uma pessoa e imprima seus dados.

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import "fmt"

type Pessoa struct {
    Nome  string
    Idade int
}

func main() {
    p := Pessoa{Nome: "Gabriel", Idade: 25}
    fmt.Printf("%s tem %d anos.\n", p.Nome, p.Idade)
}
```

**Por quê:** `type X struct {}` define um novo tipo. Usar os nomes dos campos
(`Nome:`, `Idade:`) deixa claro e protege contra erros se a ordem mudar.
</details>

---

## Exercício 2: Método Apresentar

Adicione um método `Apresentar()` à struct `Pessoa` que retorna uma string
no formato `"Oi, sou <nome> e tenho <idade> anos."`.

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import "fmt"

type Pessoa struct {
    Nome  string
    Idade int
}

// (p Pessoa) é o "receiver": liga o método ao tipo Pessoa
func (p Pessoa) Apresentar() string {
    return fmt.Sprintf("Oi, sou %s e tenho %d anos.", p.Nome, p.Idade)
}

func main() {
    p := Pessoa{Nome: "Gabriel", Idade: 25}
    fmt.Println(p.Apresentar())
}
```

**Por quê:** o `(p Pessoa)` antes do nome do método é o receiver. É o que
diferencia um método de uma função comum. `Sprintf` monta a string sem imprimir.
</details>

---

## Exercício 3: Valor vs Ponteiro (o pulo do gato)

Crie um método `Aniversario()` que aumenta a idade em 1. Faça funcionar de
verdade (a idade precisa mudar no objeto original).

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import "fmt"

type Pessoa struct {
    Nome  string
    Idade int
}

// receiver por PONTEIRO (*Pessoa): altera o objeto original
func (p *Pessoa) Aniversario() {
    p.Idade++
}

func main() {
    p := Pessoa{Nome: "Gabriel", Idade: 25}
    p.Aniversario()
    fmt.Println(p.Idade) // 26
}
```

**Por quê:** se usássemos `(p Pessoa)` (por valor), o método receberia uma
**cópia** e a mudança se perderia. Com `(p *Pessoa)` (ponteiro), alteramos
o original. Regra prática: **se o método modifica o struct, use ponteiro.**
</details>

---

## Exercício 4: Composição (Go não tem herança)

Crie uma struct `Funcionario` que "tem" uma `Pessoa` embutida e adiciona
`Cargo`. Acesse o nome do funcionário diretamente.

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import "fmt"

type Pessoa struct {
    Nome  string
    Idade int
}

type Funcionario struct {
    Pessoa // campo embutido (embedding)
    Cargo  string
}

func main() {
    f := Funcionario{
        Pessoa: Pessoa{Nome: "Gabriel", Idade: 25},
        Cargo:  "Arquiteto de Soluções",
    }
    // graças ao embedding, acessamos f.Nome direto, sem f.Pessoa.Nome
    fmt.Printf("%s — %s\n", f.Nome, f.Cargo)
}
```

**Por quê:** Go não tem herança. No lugar, usa **composição por embedding**:
ao colocar `Pessoa` sem nome de campo, o `Funcionario` "herda" os campos e
métodos dela. É a filosofia "composição sobre herança" embutida na linguagem.
</details>

---

## Exercício 5: Conta bancária

Crie uma struct `Conta` com `Saldo float64` e métodos `Depositar(valor)` e
`Sacar(valor)`. O saque não pode deixar o saldo negativo.

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import (
    "errors"
    "fmt"
)

type Conta struct {
    Saldo float64
}

func (c *Conta) Depositar(valor float64) {
    c.Saldo += valor
}

func (c *Conta) Sacar(valor float64) error {
    if valor > c.Saldo {
        return errors.New("saldo insuficiente")
    }
    c.Saldo -= valor
    return nil
}

func main() {
    c := Conta{Saldo: 100}
    c.Depositar(50)            // saldo 150
    if err := c.Sacar(200); err != nil {
        fmt.Println("Erro:", err) // saldo insuficiente
    }
    c.Sacar(30) // saldo 120
    fmt.Printf("Saldo final: %.2f\n", c.Saldo)
}
```

**Por quê:** ambos os métodos usam ponteiro porque alteram o saldo. O `Sacar`
retorna `error` — o jeito idiomático de Go sinalizar falha (veremos a fundo
na Etapa 05). Note o padrão `if err := ...; err != nil` muito comum em Go.
</details>

---

⬅️ Anterior: [Etapa 01](../01-fundamentos/exercicios.md) · ➡️ Próxima: [Etapa 03 — Interfaces](../03-interfaces/exercicios.md)
