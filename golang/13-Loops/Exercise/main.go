package main

import "fmt"

func main() {
	var day uint
	for {
		fmt.Println("Digite um número de 0-7: ")
		fmt.Scan(&day)

		if day == 0 {
			fmt.Println("Parabens! Você escolheu o número 0")
			break
		}

		switch day {
		case 1:
			fmt.Println("Sabado")
		case 2:
			fmt.Println("Domingo")
		case 3:
			fmt.Println("Ninguem gosta de segunda, vamos logo para terça XD")
			fallthrough
		case 4:
			fmt.Println("Terça")
		case 5:
			fmt.Println("Quarta")
		case 6:
			fmt.Println("Quinta")
		case 7:
			fmt.Println("Sexta")
		default:
			fmt.Println("Valor Invalido")
		}
	}
}
