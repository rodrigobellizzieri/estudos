package main

import "fmt"

type person struct {
	name string
	lastName	string
	age	uint8
}

type student struct {
	person
	college string
	course	string
}

func main() {
	p1 := person{"Junior", "Cardoso", 28}
	fmt.Println(p1)

	e1 := student{p1, "Anhembi Morumbi", "Design"}
	fmt.Println(e1)

	fmt.Println(e1.lastName)
}