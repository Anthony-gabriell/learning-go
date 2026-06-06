# Etapa 06: APIs HTTP

Onde Go brilha de verdade no mercado. A stdlib (`net/http`, `encoding/json`)
já entrega tudo para construir APIs robustas sem framework.

---

## Exercício 1: Servidor "Hello"

Suba um servidor HTTP na porta 8080 que responde "Olá, mundo!" na rota `/`.

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Olá, mundo!")
    })
    fmt.Println("Servidor em http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
```

Teste no navegador ou com `curl http://localhost:8080`.

**Por quê:** `HandleFunc` liga uma rota a uma função. `w` é onde escrevemos a
resposta, `r` traz os dados da requisição. `ListenAndServe` sobe o servidor.
Tudo isso na biblioteca padrão.
</details>

---

## Exercício 2: Respondendo JSON

Crie uma rota `/usuario` que retorna um JSON `{"nome":"Gabriel","idade":25}`.

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import (
    "encoding/json"
    "net/http"
)

type Usuario struct {
    Nome  string `json:"nome"`
    Idade int    `json:"idade"`
}

func main() {
    http.HandleFunc("/usuario", func(w http.ResponseWriter, r *http.Request) {
        u := Usuario{Nome: "Gabriel", Idade: 25}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(u)
    })
    http.ListenAndServe(":8080", nil)
}
```

**Por quê:** as tags `json:"nome"` controlam como o campo aparece no JSON
(minúsculo, como manda a convenção). `json.NewEncoder(w).Encode(u)` serializa
o struct direto na resposta. Setar o `Content-Type` é boa prática.
</details>

---

## Exercício 3: Lendo JSON do corpo (POST)

Crie uma rota `/usuario` que aceita POST, lê um JSON do corpo e devolve uma
saudação com o nome recebido.

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Usuario struct {
    Nome string `json:"nome"`
}

func main() {
    http.HandleFunc("/usuario", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "use POST", http.StatusMethodNotAllowed)
            return
        }
        var u Usuario
        if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
            http.Error(w, "JSON inválido", http.StatusBadRequest)
            return
        }
        fmt.Fprintf(w, "Olá, %s!", u.Nome)
    })
    http.ListenAndServe(":8080", nil)
}
```

Teste:
```bash
curl -X POST http://localhost:8080/usuario -d '{"nome":"Gabriel"}'
```

**Por quê:** `json.NewDecoder(r.Body).Decode(&u)` lê e converte o corpo no
struct (note o `&` — passamos o ponteiro para ser preenchido). Sempre cheque
o método HTTP e trate JSON malformado com o status code certo.
</details>

---

## Exercício 4: Status codes e validação

Na rota do exercício anterior, retorne `400 Bad Request` se o nome vier vazio.

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Usuario struct {
    Nome string `json:"nome"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    var u Usuario
    if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
        http.Error(w, "JSON inválido", http.StatusBadRequest)
        return
    }
    if u.Nome == "" {
        http.Error(w, "nome é obrigatório", http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusCreated) // 201
    fmt.Fprintf(w, "Usuário %s criado!", u.Nome)
}

func main() {
    http.HandleFunc("/usuario", handler)
    http.ListenAndServe(":8080", nil)
}
```

**Por quê:** status codes corretos são parte de uma boa API. `http.Error` já
escreve a mensagem e o código de erro. `w.WriteHeader(http.StatusCreated)` define
201 para criação bem-sucedida. As constantes (`http.StatusBadRequest` etc.)
deixam o código legível.
</details>

---

## Exercício 5: Múltiplas rotas com ServeMux

Organize três rotas (`/`, `/sobre`, `/contato`) usando um `http.ServeMux`
explícito em vez do mux global.

<details>
<summary>👉 Ver resposta</summary>

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Página inicial")
    })
    mux.HandleFunc("/sobre", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Sobre nós")
    })
    mux.HandleFunc("/contato", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Fale conosco")
    })

    fmt.Println("Servidor em http://localhost:8080")
    http.ListenAndServe(":8080", mux) // passamos o mux aqui
}
```

**Por quê:** criar um `ServeMux` explícito (em vez do global padrão) é mais
limpo, testável e evita conflitos em projetos maiores. É o passo natural antes
de partir para roteadores como `chi` ou `gorilla/mux` em projetos sérios.
</details>

---

⬅️ Anterior: [Etapa 05](../05-erros-testes/exercicios.md) · ➡️ Próxima: [Etapa 07 — Projeto Final](../07-projeto-final/exercicios.md)
