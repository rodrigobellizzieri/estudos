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

	// Array Internos
	fmt.Println("-------------")
	slice3 := make([]int32, 10, 15)
	fmt.Println(slice3)
	fmt.Println("Tamanho:", len(slice3))
	fmt.Println("Capacidade:", cap(slice3))
	slice3 = append(slice3, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	slice3[1] = 2
	fmt.Println(slice3)
	fmt.Println("Tamanho:", len(slice3))
	fmt.Println("Capacidade:", cap(slice3))

}
