package main

import "fmt"

func normalFuntion(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFuntion("Hello")
	tripleArgument(1, 2, "Alvaro")
	fmt.Println(returnValue(2))

	value1, value2 := doubleReturn(2)
	// value1, _ := doubleReturn(2) - descarta segundo valor
	fmt.Println("Value1:", value1, "Value2:", value2)
}
