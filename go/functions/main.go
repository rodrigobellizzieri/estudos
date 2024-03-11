package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 234, 564, 342, 657, 7762, 5634, 123))
}

func sum(numeros ... int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	fmt.Println(total)
	return total
}