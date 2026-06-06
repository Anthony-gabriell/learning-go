# Etapa 05: Erros e Testes

Go não tem exceptions. Erros são **valores** que você retorna e trata
explicitamente. E os testes são nativos: `go test`, sem framework externo.

---

## Exercício 1: Retornando erro

Escreva `dividir(a, b float64) (float64, error)` que retorna erro se `b == 0`.

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import (
    "errors"
    "fmt"
)

func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("não é possível dividir por zero")
    }
    return a / b, nil
}

func main() {
    resultado, err := dividir(10, 2)
    if err != nil {
        fmt.Println("Erro:", err)
        return
    }
    fmt.Println("Resultado:", resultado) // 5
}
```

**Por quê:** o padrão de Go é retornar `(valor, error)`. Quem chama **sempre**
checa `if err != nil`. Não há try/catch — o tratamento é explícito e local.
</details>

---

## Exercício 2: Erro com contexto (fmt.Errorf)

Modifique a função para incluir os valores no texto do erro, usando `fmt.Errorf`.

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import "fmt"

func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("divisão inválida: %g / %g", a, b)
    }
    return a / b, nil
}

func main() {
    _, err := dividir(10, 0)
    if err != nil {
        fmt.Println(err) // divisão inválida: 10 / 0
    }
}
```

**Por quê:** `fmt.Errorf` cria erros com mensagem formatada. Para encadear erros
(error wrapping) usa-se o verbo `%w`, que permite recuperar o erro original
depois com `errors.Is` / `errors.As`.
</details>

---

## Exercício 3: Seu primeiro teste

Crie um arquivo `mat.go` com a função `Dobro(n int) int` e um `mat_test.go`
que a testa com `go test`.

<details>
<summary>👉 Ver resposta</summary>

**mat.go**
```go
package mat

func Dobro(n int) int {
    return n * 2
}
```

**mat_test.go**
```go
package mat

import "testing"

func TestDobro(t *testing.T) {
    resultado := Dobro(5)
    esperado := 10
    if resultado != esperado {
        t.Errorf("Dobro(5) = %d; esperado %d", resultado, esperado)
    }
}
```

Rode com:
```bash
go test ./...
```

**Por quê:** arquivos de teste terminam em `_test.go`. Funções de teste começam
com `Test` e recebem `*testing.T`. `t.Errorf` marca o teste como falho. Tudo
nativo, sem instalar nada.
</details>

---

## Exercício 4: Table-driven tests

Teste a função `Dobro` com vários casos de uma vez, usando o padrão idiomático
de "table-driven tests".

<details>
<summary>👉 Ver resposta</summary>

```go
package mat

import "testing"

func TestDobroVariosCasos(t *testing.T) {
    casos := []struct {
        nome     string
        entrada  int
        esperado int
    }{
        {"positivo", 5, 10},
        {"zero", 0, 0},
        {"negativo", -3, -6},
    }

    for _, c := range casos {
        t.Run(c.nome, func(t *testing.T) {
            if got := Dobro(c.entrada); got != c.esperado {
                t.Errorf("Dobro(%d) = %d; esperado %d", c.entrada, got, c.esperado)
            }
        })
    }
}
```

**Por quê:** esse é **o** padrão de teste em Go. Você define uma tabela (slice
de structs) com os casos e itera. `t.Run` cria um subteste nomeado para cada
caso — se um falhar, você sabe exatamente qual. Adicionar um caso novo é só
uma linha.
</details>

---

## Exercício 5: Testando erros

Escreva um teste para a função `dividir` do Exercício 1 que verifica se o erro
aparece quando deve.

<details>
<summary>👉 Ver resposta</summary>

```go
package mat

import "testing"

func Dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("divisão por zero")
    }
    return a / b, nil
}

func TestDividir(t *testing.T) {
    t.Run("divisão válida", func(t *testing.T) {
        got, err := Dividir(10, 2)
        if err != nil {
            t.Fatalf("não esperava erro, veio: %v", err)
        }
        if got != 5 {
            t.Errorf("esperado 5, veio %g", got)
        }
    })

    t.Run("divisão por zero", func(t *testing.T) {
        _, err := Dividir(10, 0)
        if err == nil {
            t.Error("esperava um erro, mas veio nil")
        }
    })
}
```
*(lembre de `import "errors"` no topo junto com testing)*

**Por quê:** testamos os dois caminhos — sucesso e falha. `t.Fatalf` para o
teste na hora (útil quando continuar não faz sentido); `t.Errorf` registra mas
segue. Testar o caminho de erro é tão importante quanto testar o de sucesso.
</details>

---

⬅️ Anterior: [Etapa 04](../04-concorrencia/exercicios.md) · ➡️ Próxima: [Etapa 06 — APIs HTTP](../06-apis-http/exercicios.md)
