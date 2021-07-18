package main

import "fmt"

func main() {
	value_1 := 1
	value_2 := 1

	if value_1 == 1 {
		fmt.Println("This is one")
	} else {
		fmt.Println("This is not one")
	}

	if value_1 == 1 && value_2 == 1 {
		fmt.Println("These are one")
	}

	if value_1 == 1 || value_2 == 1 {
		fmt.Println("There ara at least one")
	}

	numberValidation(5)
	fmt.Println(passwordValidation("alvaro", "alvaro"))

	switch mod := 5 % 2; mod {
	case 0:
		fmt.Println("is Odd")
	default:
		fmt.Println("is Even")
	}

	value := 101
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condicion")
	}
}

func numberValidation(number int) {
	if number%2 == 0 {
		fmt.Println("This is odd")
	} else {
		fmt.Println("This is even")
	}
}

func passwordValidation(user, password string) bool {
	nickname := "alvaro"
	pass := "alvaro"
	return user == nickname && password == pass
}
