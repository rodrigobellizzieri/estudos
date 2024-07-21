package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome": "Junior",
		"sobrenome": "Caetano",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["sobrenome"])


}