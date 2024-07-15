package main

import "fmt"

func main() {
	fmt.Println("COPY")

	var1 := 10
	var2 := var1
	fmt.Println(var1, var2)

	var1++
	fmt.Println(var1, var2)

	var1++
	fmt.Println(var1, var2)

	fmt.Println("PONTEIRO")

	var var3 int = 20
	var ponteiro *int = &var3
	fmt.Println(var3, *ponteiro)

	var3++
	fmt.Println(var3, *ponteiro)

	var3++
	fmt.Println(var3, *ponteiro)


}