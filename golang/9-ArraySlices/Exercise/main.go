package main

import (
	"fmt"
)


func main() {

	// Loop array
	fmt.Println("===Exercise Array===")
	myArray := [5]int{10, 20, 30, 40, 50}
	for _, value := range myArray {
		fmt.Println("Array Number: ", value)
	}

	// Loop and append Slice
	fmt.Println("===Exercise Slice===")
	slice := []int{}
	for _, value := range myArray {
		slice = append(slice, value)
		fmt.Println("Append slice: ", slice)
	}
	fmt.Println("Finish append loop: ", slice)

	slice = append(slice, 60, 70)
	fmt.Println(slice)
}