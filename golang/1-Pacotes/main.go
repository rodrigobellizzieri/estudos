package main

import (
	"fmt"
	"giropops/auxiliar"
 
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo dentro do pacote main!")
	auxiliar.Escrever()

	error := checkmail.ValidateFormat("email@gmail.com")
	fmt.Println(error)
}
