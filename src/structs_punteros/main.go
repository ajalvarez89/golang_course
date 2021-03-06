package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) ping() {
	fmt.Println(myPc.brand, "pong")
}

func (myPc *pc) duplicateRam() {
	myPc.ram = myPc.ram * 2
}

func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GB de RAM, %d GB de Disco y es una %s", myPc.ram, myPc.disk, myPc.brand)
}

func main() {
	a := 50
	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 10
	fmt.Println(a)

	myPc := pc{ram: 16, disk: 500, brand: "Asus"}
	fmt.Println(myPc)
	myPc.ping()

	fmt.Println(myPc.ram)
	myPc.duplicateRam()

	fmt.Println(myPc.ram)
	myPc.duplicateRam()

	fmt.Println(myPc.ram)
}
