package main

import "fmt"

func main() {
	//array intialization
	a := [...]int{1, 2, 3, 4, 5}
	a[2] = 0
	b := &a
	a[3] = 0

	fmt.Println(a)
	fmt.Println(b)
	//slice intilization

	sampleSlice := []int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3}
	fmt.Println(len(sampleSlice))
	fmt.Println(cap(sampleSlice))
	fmt.Println(sampleSlice)

}
