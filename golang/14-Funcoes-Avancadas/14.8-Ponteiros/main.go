package main

import "fmt"

func ponteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	fmt.Println("Funções com Ponteiros")

	valor := 10
	fmt.Println("Valor Original:", valor)
	fmt.Println("Valor na memória:", &valor)
	fmt.Println("====================")

	ponteiro(&valor)

	fmt.Println("Valor Atual:", valor)
	fmt.Println("Valor na memória:", &valor)
}