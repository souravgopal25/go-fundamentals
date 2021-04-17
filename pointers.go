package main

import "fmt"

func changeValue(str *string) {
	*str = "CHANGED!"
}

func main() {
	//Gives address/refference
	x := 7
	y := &x
	fmt.Println(*y)
	*y = 8
	fmt.Println(*y)
	fmt.Println(x)
	toChange := "Hello"
	fmt.Println(toChange)
	changeValue(&toChange)
	fmt.Println(toChange)

}
