package main

import "fmt"

func calculos(n1, n2 int) (soma, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2

	return
}

func main() {
	soma, subtracao := calculos(20, 10)
	
	fmt.Println("Resultado Soma:", soma)
	fmt.Println("Resultado Subtração:", subtracao)
}