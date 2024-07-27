package main

import "fmt"

func recuperar() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func media(n1, n2 int) int {
	defer recuperar()
	
	resultado := (n1 + n2) / 2
	
	if resultado > 6 {
		return resultado
	} else if resultado < 6 {
		return resultado
	}

	panic("A NOTA É IGUAL A 6!!!!!!!!")
}

func main() {
	fmt.Println(media(6, 6))
}
