package main

import (
	"fmt"
	// "time"
)

func main() {
	// total := 0

	// for total < 10 {
	// 	total++
	// 	fmt.Println("Valor total:", total)
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println(total)

	// nomes := [5]string{"Gabriel", "Gustavo", "Murilo", "Eduardo", "Raphael"}

	// for indice, valor := range nomes {
	// 	fmt.Println("Indice:",indice)
	// 	fmt.Println("Nome:",valor)
	// 	time.Sleep(time.Second)
	// }

	for indice, letra := range "RODRIGO" {
		// fmt.Println(indice, letra)
		fmt.Println(indice, string(letra))
	}
}