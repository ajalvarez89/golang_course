package main

import (
	"fmt"
	"math"
)

func main() {

	// Area cuadrado
	const baseCuadrado = 100
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

	x := 5
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma:", result)

	// Resta
	result = y - x
	fmt.Println("Resta:", result)

	// Multiplicacion
	result = y * x
	fmt.Println("Multiplicacion:", result)

	// Division
	result = y / x
	fmt.Println("Division:", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo:", result)

	// Incremental
	x++
	fmt.Println("Incremental:", x)

	// Decremental
	x--
	fmt.Println("Decremental:", x)

	// Area rectangulo
	const baseRectangulo = 50
	const alturaRectangulo = 30
	areaRectangulo := baseRectangulo * alturaRectangulo
	fmt.Println("Area rectangulo:", areaRectangulo)

	// Area circulo
	var radioCirulo float64 = 30
	areaCirculo := math.Pi * math.Pow(radioCirulo, 2)
	fmt.Println("Area circulo:", areaCirculo)
}
