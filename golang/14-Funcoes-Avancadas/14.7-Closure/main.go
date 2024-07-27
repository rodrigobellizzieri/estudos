package main

import "fmt"

func testeClosure() func() {
	texto := "Estou na função Closure"

	funcao := func ()  {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Estou na função main"
	fmt.Println(texto)

	outraFuncao := testeClosure()

	outraFuncao()

}