package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	var array1 [3]string
	array1[1] = "Posição 2"
	fmt.Println(array1)
	array1[0] = "Posição 1"
	fmt.Println(array1)
	array1[0] = "Posição 3"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	fmt.Println("Slices")

	slice1 := []int{1, 2, 3, 4}
	fmt.Println(slice1)

	slice1 = append(slice1, 5)
	fmt.Println(slice1)

	slice2 := array2[2:4]
	fmt.Println(slice2)

}
