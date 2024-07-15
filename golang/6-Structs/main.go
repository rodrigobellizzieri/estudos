package main

import "fmt"

type user struct {
	name  string
	email string
	idade uint8
}

func main() {
	fmt.Println("Falando sobre Structs")

	var usuario user
	usuario.name = "Joãozinho"
	usuario.email = "joaozinho@gmail.com"
	usuario.idade = 18

	fmt.Println(usuario)

	usuario2 := user{"Pedrinho", "pedrinho@gmail.com", 23}
	fmt.Println(usuario2)
}
