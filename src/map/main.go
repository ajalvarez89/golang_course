package main

import "fmt"

func main() {
	m := make(map[string]string)

	m["name"] = "alvaro"
	m["age"] = "31"

	// Recorrer el map
	for key, value := range m {
		fmt.Println(key, value)
	}

	// Encontrar Valores
	value, ok := m["name"]
	fmt.Println(value, ok)
}
