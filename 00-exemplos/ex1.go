// ============================================================
// EXEMPLO 1 — A SIMPLICIDADE DO GO
// Um servidor web COMPLETO e funcional em ~15 linhas.
// Sem framework, sem dependência externa: só a stdlib.
// ============================================================

package main

import (
	"fmt"      // formatação de texto
	"net/http" // servidor HTTP nativo — já vem na linguagem!
)

// handler é a função que responde cada requisição.
// w = onde escrevemos a resposta | r = dados da requisição.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá! Servidor Go rodando 🐹")
}

func main() {
	// Liga a rota "/" à nossa função handler.
	http.HandleFunc("/", handler)

	fmt.Println("Servidor no ar em http://localhost:8080")

	// Sobe o servidor na porta 8080.
	// Em Java/Node você precisaria de bibliotecas extras pra isso.
	http.ListenAndServe(":8080", nil)
}

// ============================================================
// PRA RODAR:
//   go run exemplo1_servidor.go
// e abra http://localhost:8080 no navegador.
// ============================================================
