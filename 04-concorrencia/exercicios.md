# Etapa 04: Concorrência

O diferencial do Go. Goroutines (tarefas leves em paralelo) e channels
(comunicação entre elas). O lema da linguagem: *"Don't communicate by sharing
memory; share memory by communicating."*

---

## Exercício 1: Primeira goroutine

Dispare uma goroutine que imprime "Olá da goroutine!". Garanta que o programa
não termine antes dela rodar.

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    go func() {
        fmt.Println("Olá da goroutine!")
    }()

    time.Sleep(100 * time.Millisecond) // dá tempo da goroutine rodar
    fmt.Println("Fim do main")
}
```

**Por quê:** `go` dispara a função em paralelo. O problema: o `main` não espera
a goroutine. Se ele terminar antes, a goroutine morre junto. O `Sleep` resolve
de forma tosca — a forma correta vem no próximo exercício (WaitGroup).
</details>

---

## Exercício 2: WaitGroup (esperar do jeito certo)

Dispare 3 goroutines e espere TODAS terminarem antes de prosseguir, sem usar
`time.Sleep`.

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1) // registra 1 tarefa
        go func(id int) {
            defer wg.Done() // avisa que terminou
            fmt.Printf("Goroutine %d concluída\n", id)
        }(i) // passamos i como argumento (importante!)
    }

    wg.Wait() // bloqueia até todas chamarem Done()
    fmt.Println("Todas terminaram")
}
```

**Por quê:** `WaitGroup` é um contador. `Add(1)` incrementa, `Done()` decrementa,
`Wait()` bloqueia até zerar. **Detalhe crucial:** passamos `i` como argumento
`(i)` — se usássemos `i` direto dentro da closure, todas poderiam pegar o mesmo
valor final (clássico bug de concorrência).
</details>

---

## Exercício 3: Channel básico

Crie uma goroutine que calcula a soma de um slice e envia o resultado de volta
ao `main` por um channel.

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import "fmt"

func somar(nums []int, resultado chan int) {
    soma := 0
    for _, n := range nums {
        soma += n
    }
    resultado <- soma // envia para o channel
}

func main() {
    nums := []int{1, 2, 3, 4, 5}
    resultado := make(chan int) // cria o channel

    go somar(nums, resultado)

    total := <-resultado // recebe (bloqueia até chegar algo)
    fmt.Println("Soma:", total) // 15
}
```

**Por quê:** channel é um "cano" tipado entre goroutines. `<-` envia ou recebe.
Receber de um channel **bloqueia** até ter algo — por isso não precisamos de
WaitGroup aqui: o `<-resultado` já espera naturalmente.
</details>

---

## Exercício 4: Worker pool

Distribua 5 tarefas entre 3 workers usando channels. Cada worker pega uma
tarefa, "processa" (multiplica por 2) e devolve.

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import (
    "fmt"
    "sync"
)

func worker(id int, tarefas <-chan int, resultados chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for t := range tarefas { // consome até o channel fechar
        resultados <- t * 2
    }
}

func main() {
    tarefas := make(chan int, 5)
    resultados := make(chan int, 5)
    var wg sync.WaitGroup

    // sobe 3 workers
    for w := 1; w <= 3; w++ {
        wg.Add(1)
        go worker(w, tarefas, resultados, &wg)
    }

    // envia 5 tarefas
    for i := 1; i <= 5; i++ {
        tarefas <- i
    }
    close(tarefas) // sinaliza que não vem mais tarefa

    wg.Wait()
    close(resultados)

    for r := range resultados {
        fmt.Print(r, " ") // 2 4 6 8 10 (ordem pode variar)
    }
    fmt.Println()
}
```

**Por quê:** padrão clássico de Go. `<-chan int` é channel só de leitura,
`chan<- int` só de escrita (segurança de tipo). `close()` avisa o `range` para
parar. Workers pegam tarefas conforme ficam livres — paralelismo real.
</details>

---

## Exercício 5: Select com timeout

Use `select` para receber de um channel, mas desistir se demorar mais de
1 segundo.

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan string)

    go func() {
        time.Sleep(2 * time.Second) // simula tarefa lenta
        ch <- "resultado"
    }()

    select {
    case msg := <-ch:
        fmt.Println("Recebido:", msg)
    case <-time.After(1 * time.Second):
        fmt.Println("Timeout! Demorou demais.")
    }
}
```

**Por quê:** `select` espera em vários channels ao mesmo tempo e age no primeiro
que ficar pronto. `time.After` devolve um channel que "dispara" após o tempo.
Como a tarefa leva 2s e o timeout é 1s, o timeout vence. Esse padrão é essencial
em chamadas de rede/API que podem travar.
</details>

---

⬅️ Anterior: [Etapa 03](../03-interfaces/exercicios.md) · ➡️ Próxima: [Etapa 05 — Erros e Testes](../05-erros-testes/exercicios.md)
