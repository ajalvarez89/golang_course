package main

import "fmt"

func main() {
	name := "Alvaro"
	age := 31

	// Printf
	fmt.Printf("%v tiene %v Años\n", name, age)
	fmt.Printf("%s tiene %d Años\n", name, age)

	//Sprintf
	message := fmt.Sprintf("%s tiene %d Años", name, age)
	fmt.Println(message)

	// Tipo de datos
	fmt.Printf("name: %T\n", name)
}
