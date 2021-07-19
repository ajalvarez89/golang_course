package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Toyota", year: 2020}
	fmt.Println(myCar.brand)
	fmt.Println(myCar.year)

	var otherCar car
	otherCar.brand = "Mazda"
	fmt.Println(otherCar)

}
