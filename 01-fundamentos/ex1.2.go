package main

import "fmt"

var y = "Eren" // para que fosse compilado, declaramos var fora do bloco
var z = "Levi"

func main() {
	x := "Mikasa"
	//	y := "Eren" Foi passado a variavel dentro do bloco mas a funcao fora nao pode acessar
	exibirnomes(x) // chamamos essa funcao e passamos x
}

func exibirnomes(x string) { // Precisa especificar o tipo-primitivo que o x vai receber
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
