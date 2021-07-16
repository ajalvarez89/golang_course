package main

import "fmt"

func main() {
	// Constantes
	const pi float64 = 3.1415926535
	const pi2 = 3.15

	fmt.Println("pi: ", pi)
	fmt.Println("***************")
	fmt.Println("pi2: ", pi2)

	// Enteros
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 100
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Area cuadrado:", areaCuadrado)
}
