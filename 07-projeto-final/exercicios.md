# Etapa 07: Projeto Final

Hora de juntar tudo. Você vai construir uma **API REST de Tarefas (To-Do)**
em memória, usando structs, métodos, interfaces, tratamento de erro, JSON e HTTP.

Em vez de exercícios soltos, aqui é **um projeto incremental**. Tente fazer
cada etapa antes de olhar a referência.

---

## O desafio

Construa uma API com os seguintes endpoints:

| Método | Rota          | O que faz                  |
|--------|---------------|----------------------------|
| GET    | `/tarefas`    | Lista todas as tarefas     |
| POST   | `/tarefas`    | Cria uma nova tarefa       |
| PUT    | `/tarefas/concluir?id=N` | Marca tarefa como concluída |

Cada tarefa tem: `ID`, `Titulo` e `Concluida`.

---

## Etapa A: Modelo e armazenamento

Crie a struct `Tarefa` e um "repositório" em memória (um slice + um contador
de ID) com métodos para listar, adicionar e concluir.

<details>
<summary>👉 Ver referência</summary>

```go
type Tarefa struct {
    ID        int    `json:"id"`
    Titulo    string `json:"titulo"`
    Concluida bool   `json:"concluida"`
}

type Repositorio struct {
    tarefas []Tarefa
    proxID  int
    mu      sync.Mutex // protege contra acesso concorrente
}

func NovoRepositorio() *Repositorio {
    return &Repositorio{proxID: 1}
}

func (r *Repositorio) Listar() []Tarefa {
    r.mu.Lock()
    defer r.mu.Unlock()
    return r.tarefas
}

func (r *Repositorio) Adicionar(titulo string) Tarefa {
    r.mu.Lock()
    defer r.mu.Unlock()
    t := Tarefa{ID: r.proxID, Titulo: titulo}
    r.proxID++
    r.tarefas = append(r.tarefas, t)
    return t
}

func (r *Repositorio) Concluir(id int) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    for i := range r.tarefas {
        if r.tarefas[i].ID == id {
            r.tarefas[i].Concluida = true
            return nil
        }
    }
    return fmt.Errorf("tarefa %d não encontrada", id)
}
```

**Por quê:** o `sync.Mutex` é essencial — um servidor HTTP atende requisições
em goroutines simultâneas, então o slice compartilhado precisa de proteção
contra acesso concorrente. Repare em `for i := range` (com índice) para poder
**alterar** o elemento no slice.
</details>

---

## Etapa B: Handlers HTTP

Crie os handlers que conectam as rotas ao repositório, lendo/escrevendo JSON.

<details>
<summary>👉 Ver referência</summary>

```go
func (r *Repositorio) handlerTarefas(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    switch req.Method {
    case http.MethodGet:
        json.NewEncoder(w).Encode(r.Listar())

    case http.MethodPost:
        var entrada struct {
            Titulo string `json:"titulo"`
        }
        if err := json.NewDecoder(req.Body).Decode(&entrada); err != nil {
            http.Error(w, "JSON inválido", http.StatusBadRequest)
            return
        }
        if entrada.Titulo == "" {
            http.Error(w, "título é obrigatório", http.StatusBadRequest)
            return
        }
        nova := r.Adicionar(entrada.Titulo)
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(nova)

    default:
        http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
    }
}

func (r *Repositorio) handlerConcluir(w http.ResponseWriter, req *http.Request) {
    idStr := req.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "id inválido", http.StatusBadRequest)
        return
    }
    if err := r.Concluir(id); err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "tarefa %d concluída", id)
}
```

**Por quê:** um handler por recurso, com `switch req.Method` separando GET/POST.
`req.URL.Query().Get("id")` lê query params; `strconv.Atoi` converte string em
int (com erro tratado). Os métodos do repositório ficam ligados a ele via receiver.
</details>

---

## Etapa C: Juntando no main

Monte o `main` que cria o repositório, registra as rotas e sobe o servidor.

<details>
<summary>👉 Ver referência — programa completo</summary>

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
    "sync"
)

type Tarefa struct {
    ID        int    `json:"id"`
    Titulo    string `json:"titulo"`
    Concluida bool   `json:"concluida"`
}

type Repositorio struct {
    tarefas []Tarefa
    proxID  int
    mu      sync.Mutex
}

func NovoRepositorio() *Repositorio {
    return &Repositorio{proxID: 1}
}

func (r *Repositorio) Listar() []Tarefa {
    r.mu.Lock()
    defer r.mu.Unlock()
    return r.tarefas
}

func (r *Repositorio) Adicionar(titulo string) Tarefa {
    r.mu.Lock()
    defer r.mu.Unlock()
    t := Tarefa{ID: r.proxID, Titulo: titulo}
    r.proxID++
    r.tarefas = append(r.tarefas, t)
    return t
}

func (r *Repositorio) Concluir(id int) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    for i := range r.tarefas {
        if r.tarefas[i].ID == id {
            r.tarefas[i].Concluida = true
            return nil
        }
    }
    return fmt.Errorf("tarefa %d não encontrada", id)
}

func (r *Repositorio) handlerTarefas(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    switch req.Method {
    case http.MethodGet:
        json.NewEncoder(w).Encode(r.Listar())
    case http.MethodPost:
        var entrada struct {
            Titulo string `json:"titulo"`
        }
        if err := json.NewDecoder(req.Body).Decode(&entrada); err != nil {
            http.Error(w, "JSON inválido", http.StatusBadRequest)
            return
        }
        if entrada.Titulo == "" {
            http.Error(w, "título é obrigatório", http.StatusBadRequest)
            return
        }
        nova := r.Adicionar(entrada.Titulo)
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(nova)
    default:
        http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
    }
}

func (r *Repositorio) handlerConcluir(w http.ResponseWriter, req *http.Request) {
    id, err := strconv.Atoi(req.URL.Query().Get("id"))
    if err != nil {
        http.Error(w, "id inválido", http.StatusBadRequest)
        return
    }
    if err := r.Concluir(id); err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    fmt.Fprintf(w, "tarefa %d concluída", id)
}

func main() {
    repo := NovoRepositorio()
    mux := http.NewServeMux()
    mux.HandleFunc("/tarefas", repo.handlerTarefas)
    mux.HandleFunc("/tarefas/concluir", repo.handlerConcluir)

    fmt.Println("API de Tarefas em http://localhost:8080")
    http.ListenAndServe(":8080", mux)
}
```

### Como testar
```bash
go run main.go

# em outro terminal:
curl -X POST localhost:8080/tarefas -d '{"titulo":"Estudar Go"}'
curl -X POST localhost:8080/tarefas -d '{"titulo":"Gravar TikTok"}'
curl localhost:8080/tarefas
curl -X PUT "localhost:8080/tarefas/concluir?id=1"
curl localhost:8080/tarefas
```

**Por quê:** este projeto reúne toda a trilha — struct (modelo), métodos com
ponteiro (repositório), mutex (concorrência), JSON (encoding), HTTP (handlers)
e tratamento de erro. É um esqueleto realista de microsserviço em Go.
</details>

---

## 🚀 Próximos passos (desafios extras)

Quando terminar, evolua o projeto:

- [ ] Adicionar endpoint DELETE para remover tarefas
- [ ] Escrever testes (`_test.go`) para o repositório
- [ ] Trocar o slice por um `map[int]Tarefa` (busca O(1))
- [ ] Persistir em arquivo JSON ou num banco (SQLite)
- [ ] Trocar o mux nativo por um roteador como `chi`
- [ ] Dockerizar a aplicação (`Dockerfile`)

---

⬅️ Anterior: [Etapa 06](../06-apis-http/exercicios.md) · 🏠 [Voltar ao início](../README.md)

---

*Parabéns por chegar até aqui! Você passou pela base, concorrência e construção
de API — o núcleo do que se usa com Go em produção. 🐹*
