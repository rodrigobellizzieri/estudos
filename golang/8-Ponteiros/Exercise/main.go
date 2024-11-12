package main

import "fmt"

func increase(value *int) {
	*value++
	fmt.Println("Increase number: ", *value)
} 


func main() {
	number := 10
	fmt.Println("Start number: ", number)

	increase(&number)
}