package main

import ("fmt")

// type ID float64

// var (
// 	nome string = "Rodrigo"
// 	idade int 	= 10
// 	valor ID 	= 2.50
// 	check bool 	= true
// )

// const verdade string = "Isso é"

// // func main() {
// // 	println("Hello World!")
// // }

// func main() {
// 	fmt.Printf("Variavel nome tem o tipo: %T", nome)
// 	fmt.Printf("Variavel idade tem o tipo: %T", idade)
// 	fmt.Printf("Variavel valor tem o tipo: %T", valor)
// 	fmt.Printf("Variavel check tem o tipo: %T e valor: %v", check, check)
// 	fmt.Printf(verdade)
// }

func main() {
	var meuArray [5]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30
	meuArray[3] = 40
	meuArray[4] = 50

	for i, v := range meuArray {
		fmt.Printf("Meu indice é: %d = %d \n",i, v)
	}
}