package main

import "fmt"

func printPrice(value string) {

	fruitsPrice := map[string]float64{
		"Uva":     10.50,
		"Maçã":    5.50,
		"Laranja": 6.99,
		"Banana":  7.99,
	}

	fmt.Println("Fruit:", value)
	fmt.Println("Price:", fruitsPrice[value])
}

func main() {

	printPrice("Maçã")

}
