package main

import "fmt"

func calculoMedia(n1, n2 float32) bool {
	media := (n1 + n2) / 2
	defer fmt.Print("Status Aprovado: ")

	if media >= 6 {
		return true
	}

	
	return false
	
}

func main() {
	fmt.Println(calculoMedia(4, 7))
}