package main

import "fmt"

func somar(n1, n2 int) int {
	return n1 + n2
}

func calculos(n3, n4 int) (int, int) {
	soma := n3 + n4
	subtracao := n3 - n4
	return  soma, subtracao
	
}

func main() {
	soma := somar(10, 30)
	fmt.Println("O resultado é:",soma)

	var f = func(text string) string {
		return text

	}

	resultado := f("Hello World")
	fmt.Println(resultado)

	resultadoSoma, _ := calculos(30, 20)
	fmt.Println(resultadoSoma)

	_, resultadoSubt := calculos(30, 20)
	fmt.Println(resultadoSubt)
}