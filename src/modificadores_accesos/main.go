package main

// go mod init

import (
	pk "curso_golang/src/modificadores_accesos/mypackage"
	"fmt"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	fmt.Println(myCar)

	otherCar := pk.CarPublic{Brand: "Toyota", Year: 2021}
	fmt.Println(otherCar)

	pk.PrintMessage("This is a message")
}
