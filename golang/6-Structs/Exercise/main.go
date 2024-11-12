package main

import "fmt"

type Products struct {
	Name 		string
	Category	string
	Price		float64
	Available	bool
}


func main() {
	laptop := Products{"Dell G15", "Laptop", 3500.00, true }

	showProduct(laptop)
	descount(laptop)
}

func showProduct(p Products) {
	fmt.Println("Product info:")

	fmt.Println("Name: ", p.Name)
	fmt.Println("Category: ", p.Category)
	fmt.Println("Price: ", p.Price)
	fmt.Println("Available: ", p.Available)
}

func descount(price Products) {
	fmt.Println("==========")
	fmt.Println("Descount received 10%: ", price.Price * 0.10)
	price.Price = price.Price - (price.Price * 0.10)
	fmt.Println("New price: ", price.Price)
}