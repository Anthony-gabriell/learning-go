// ============================================================
// EXEMPLO 2 — O SUPERPODER DO GO: CONCORRÊNCIA
// Dispara milhares de tarefas em paralelo com a palavra "go".
// Goroutines + Channels = o coração do Golang.
// ============================================================

package main

import (
	"fmt"
	"sync"
	"time"
)

// tarefa simula um trabalho que leva um tempinho
// (ex: chamar uma API, ler um arquivo, consultar um banco).
func tarefa(id int, wg *sync.WaitGroup) {
	defer wg.Done() // avisa que terminou quando a função acabar

	fmt.Printf("Tarefa %d começou...\n", id)
	time.Sleep(time.Second) // simula 1s de trabalho
	fmt.Printf("Tarefa %d terminou ✅\n", id)
}

func main() {
	inicio := time.Now()

	// WaitGroup espera todas as goroutines terminarem.
	var wg sync.WaitGroup

	// Dispara 5 tarefas AO MESMO TEMPO.
	// Repare na palavrinha mágica: "go"
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go tarefa(i, &wg) // <- roda em paralelo, não trava o loop
	}

	wg.Wait() // espera todas terminarem

	// 5 tarefas de 1s cada, mas o total é ~1s (rodaram juntas!)
	fmt.Printf("\nTudo pronto em %v 🚀\n", time.Since(inicio))
}

// ============================================================
// PRA RODAR:
//   go run exemplo2_goroutines.go
//
// 💡 SACADA PRO VÍDEO:
// 5 tarefas de 1 segundo cada deveriam levar 5s no modo
// tradicional (uma após a outra). Com goroutines elas rodam
// JUNTAS e o total fica em ~1 segundo. Esse é o pulo do gato.
// ============================================================
