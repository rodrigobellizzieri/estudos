package main

import (
	"fmt"
	"time"
)

func soma(numeros ...int) {
	total := 0
	for _, numero := range numeros {
		total += numero
		time.Sleep(time.Second)
		fmt.Println("Total atual:", total)
	}

	fmt.Println("=============")
	fmt.Println("Total:",total)
}

func main() {
	soma(1, 2, 3, 4, 5, 6, 7, 8, 10, 23, 12, 545, 23, 56, 106)
}
