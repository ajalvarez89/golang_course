package main

import "fmt"

type figuras2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (c rectangulo) area() float64 {
	return c.base * c.altura
}

func calculate(f figuras2D) {
	fmt.Println("Area:", f.area())
}

func main() {
	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 2, altura: 6}

	calculate(myCuadrado)
	calculate(myRectangulo)

	// Lista de interfaces
	myInterfaces := []interface{}{"Hola", 10, 40.5}
	fmt.Println(myInterfaces)
}
