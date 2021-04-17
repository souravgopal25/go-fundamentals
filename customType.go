package main

import "fmt"

type Name struct {
	first string
	last  string
}

func (name *Name) setName(first string) {
	name.first = first

}

func main() {
	var person1 Name = Name{"Sourav", "Sharma"}
	fmt.Println(person1.first)
}
