package main

import (
	"errors"
	"fmt"
)

func main() {
	// My system 64, so int64
	var numero1 int = 10
	fmt.Println(numero1)

	var numero2 int32 = -1000000000
	fmt.Println(numero2)

	var numero3 uint = 100000
	fmt.Println(numero3)

	// My system 64, so int64
	numero4 := 10
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.90
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123455690923023804.80
	fmt.Println(numeroReal2)

	// Char
	anyValue := 'R'
	fmt.Println(anyValue)

	var booleano bool
	fmt.Println(booleano)

	var myError error = errors.New("My Internal Error!")
	fmt.Println(myError)

}
