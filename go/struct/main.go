package main

import (
	"fmt"
)

type Address struct {
	Street 	string
	Code	int
}

type Client struct {
	Name 	string
	Age 	int
	Active	bool
	Address
}

func (c Client) DisableClient() {
	c.Active = false
	fmt.Printf("O cliente %s foi desativado \n", c.Name)
	fmt.Printf("Status atual: %t \n", c.Active)
}


func main() {
	rodrigo := Client {
		Name: "Rodrigo",
		Age: 26,
		Active: true,
	}

	rodrigo.Street = "Rua Paraiba"
	rodrigo.Code = 6764040
	rodrigo.DisableClient()

	fmt.Printf("Status verdadeiro: %t \n", rodrigo.Active)
}