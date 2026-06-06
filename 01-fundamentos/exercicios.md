# Etapa 01: Fundamentos

Variáveis, tipos, loops, condicionais e funções. A base de tudo.

> Tente resolver cada exercício num arquivo `.go` antes de abrir a resposta.

---

## Exercício 1: Hello, com nome

Escreva um programa que tem uma variável com seu nome e imprime
`Olá, <nome>! Bem-vindo ao Go.`

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import "fmt"

func main() {
    nome := "Gabriel" // := declara e infere o tipo (string)
    fmt.Printf("Olá, %s! Bem-vindo ao Go.\n", nome)
}
```

**Por quê:** `:=` é a forma curta de declarar variável com inferência de tipo.
`%s` no `Printf` é o marcador para string.
</details>

---

## Exercício 2: Par ou ímpar

Escreva uma função `parOuImpar(n int) string` que retorna `"par"` ou `"ímpar"`.
Teste com alguns números no `main`.

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import "fmt"

func parOuImpar(n int) string {
    if n%2 == 0 {
        return "par"
    }
    return "ímpar"
}

func main() {
    fmt.Println(parOuImpar(4)) // par
    fmt.Println(parOuImpar(7)) // ímpar
}
```

**Por quê:** o operador `%` (módulo) dá o resto da divisão. Resto 0 = par.
Em Go não precisa de `else` aqui: se entrou no `if` e retornou, o resto não executa.
</details>

---

## Exercício 3: Soma de 1 até N

Escreva uma função `soma(n int) int` que soma todos os números de 1 até N
usando um loop `for`.

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import "fmt"

func soma(n int) int {
    total := 0
    for i := 1; i <= n; i++ {
        total += i
    }
    return total
}

func main() {
    fmt.Println(soma(5)) // 15 (1+2+3+4+5)
}
```

**Por quê:** Go tem só o `for` — não existe `while`. O formato
`for inicio; condição; incremento` é o loop clássico. Para um "while",
basta usar `for condição {}`.
</details>

---

## Exercício 4: FizzBuzz

Imprima os números de 1 a 20. Para múltiplos de 3 imprima `Fizz`,
de 5 imprima `Buzz`, e de ambos imprima `FizzBuzz`.

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 20; i++ {
        switch {
        case i%15 == 0:
            fmt.Println("FizzBuzz")
        case i%3 == 0:
            fmt.Println("Fizz")
        case i%5 == 0:
            fmt.Println("Buzz")
        default:
            fmt.Println(i)
        }
    }
}
```

**Por quê:** o `switch` sem expressão (`switch {}`) funciona como uma cadeia
de `if/else if` mais limpa. A ordem importa: testamos múltiplo de 15 primeiro
(que é múltiplo de 3 E 5 ao mesmo tempo).
</details>

---

## Exercício 5: Média de um slice

Escreva uma função `media(nums []float64) float64` que calcula a média
dos valores de um slice.

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import "fmt"

func media(nums []float64) float64 {
    if len(nums) == 0 {
        return 0 // evita divisão por zero
    }
    var soma float64
    for _, n := range nums {
        soma += n
    }
    return soma / float64(len(nums))
}

func main() {
    notas := []float64{7.5, 8.0, 6.5, 9.0}
    fmt.Printf("Média: %.2f\n", media(notas)) // 7.75
}
```

**Por quê:** `range` percorre o slice; o `_` ignora o índice (só queremos o valor).
`len()` dá o tamanho. Precisamos converter para `float64(len(...))` porque
`len` retorna `int` e Go não mistura tipos numéricos automaticamente.
</details>

---

➡️ Próxima: [Etapa 02 — Structs e Métodos](../02-structs-metodos/exercicios.md)
