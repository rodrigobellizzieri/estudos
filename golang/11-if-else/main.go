package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 10
	numero2 := 12

	if numero > 15 {
		fmt.Println("Numero maior que 15")
	} else {
		fmt.Println("Numero menor que 15")
	}

	if total := numero + numero2; total > 30 {
		fmt.Println("Total maior que 30")
	} else {
		fmt.Println("Total menor que 30")
	}

	nome := "João"

	if nome == "Julio" {
		fmt.Println("Olá Julio") 
	} else if nome == "João" {
		fmt.Println("Olá João")
	} else {
		fmt.Println("Desculpe não sei o seu nome ;(")
	}
}
