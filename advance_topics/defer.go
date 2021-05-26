package main

import "fmt"

//https://blog.golang.org/defer-panic-and-recover
//Defer is a keyword which executes after the execution of the current block and it uses stack DS
//Defer is ususally used for cleanup function
func main() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Main Function Completes")
}
