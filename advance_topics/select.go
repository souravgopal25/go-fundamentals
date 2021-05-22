package main

import "fmt"

func main() {
	c1 := make(chan int)
	c2 := make(chan string)

	go func() {
		c1 <- 5
	}()

	go func() {
		c2 <- "Hello Sourav"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1, "Received")
		case msg2 := <-c2:
			fmt.Println(msg2, "Received")

		}
	}
}
