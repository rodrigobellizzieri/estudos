package main

import (
	"github.com/badoux/checkmail"
	"fmt"
)

func main() {
	fmt.Println("Check Email Format")

	erro := checkmail.ValidateFormat("myemailgmail.com")

	if erro == nil {
		fmt.Println("Email Format - OK")
	} else {
		fmt.Println("Email Format - Invalid")
	}
}