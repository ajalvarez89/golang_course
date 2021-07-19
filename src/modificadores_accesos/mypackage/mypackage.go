package mypackage

import "fmt"

// CarPublic Car with public access
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

func PrintMessage(message string) {
	printMessage(message)
}

func printMessage(message string) {
	fmt.Println(message)
}
