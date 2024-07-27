package main

import "fmt"

type user struct {
	name 	string
	age 	uint8
}

func (u user) saveUser() {
	fmt.Println("Salvando o usuário:", u.name)
}

func main() {
	user1 := user{"Rodrigo97", 26}
	user1.saveUser()

	user2 := user{"Rafael03", 21}
	user2.saveUser()
}