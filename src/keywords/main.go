package main

import "fmt"

func main() {
	// Defer
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//Continue y break

	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//continue
		if i == 2 {
			fmt.Println("Es Dos")
			continue
		}

		//break
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}
}
