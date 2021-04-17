package main

import "fmt"

func myfunc(p, q int) (rectangle int, square int) {
	rectangle = p * q
	square = p * p
	return
}

func main() {

	var area1, area2 = myfunc(2, 4)

	// Display the values
	fmt.Printf("Area of the rectangle is: %d", area1)
	fmt.Printf("\nArea of the square is: %d", area2)

}
