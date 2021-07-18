package main

import "fmt"

func main() {
	// For condicional
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("************")

	// For While
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// For Forever
	// counterForever := 0
	// for {
	// 	fmt.Println(counterForever)
	// 	counterForever++
	// }

	// For decrementally
	for i := 10; i != 0; i-- {
		fmt.Println(i)
	}
}
