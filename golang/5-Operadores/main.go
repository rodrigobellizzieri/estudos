package main

import "fmt"

func main() {
	// ARITIMÉTICOS
	soma := 5 + 5
	subtracao := 5 -4
	divisao := 10 / 2
	multiplicacao := 5 * 5
	restodivisao := 10 % 5

	fmt.Println(soma, subtracao, divisao, multiplicacao, restodivisao)
	// FIM ARITIMÉTICOS

	// ATRIBUIÇÃO
	var variavel1 string = "Valor 1"
	variavel2 := "Valor 2"
	
	fmt.Println(variavel1, variavel2)
	// FIM ATRIBUIÇÃO

	// RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 >= 2)
	// FIM RELACIONAIS

	//LÓGICOS
	verdadeiro, falso := true, false
	fmt.Println("===========")
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	//FIM LÓGICOS

	//UNARIOS
	fmt.Println("===========")
	numero := 10
	numero++
	numero += 10
	numero --
	numero *= 2
	fmt.Println(numero)

	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)
}