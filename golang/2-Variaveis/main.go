package main

import "fmt"

func main() {
	var variavel1 string = "Hello World!"
	fmt.Println(variavel1)
	
	variavel2 := "Hello World 2!"
	fmt.Println(variavel2)

	var (
		variavel3 string = "Varaiavel 3"
		variavel4 string = "Varaiavel 4"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
