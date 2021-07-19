package main

import "fmt"

//chan<- entrada de datos
//<-chan salida de datos
func say(text string, c chan<- string) {
	c <- text
}

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 1)
	c2 := make(chan string, 2)

	fmt.Println("Hello")
	go say("Frist channel", c)
	fmt.Println(<-c)

	c2 <- "First"
	c2 <- "Second"
	//range and close
	close(c2)

	for message := range c2 {
		fmt.Println(message)
	}

	//select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("Message 1", email1)
	go message("Message 2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email Recibido de email 1", m1)
		case m2 := <-email2:
			fmt.Println("Email Recibido de email 2", m2)
		}
	}
}
