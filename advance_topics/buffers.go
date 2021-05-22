package main

import "fmt"

func main() {
	//make(type , capacity)
	channel := make(chan int, 2)
	channel <- 2
	channel <- 3
	channel <- 4
	fmt.Println(<-channel)
	fmt.Println(cap(channel))
	fmt.Println(len(channel))
}
