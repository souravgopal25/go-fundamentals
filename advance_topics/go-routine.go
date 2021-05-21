package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.golang.org",
		"https://www.twitter.com",
	}
	for _, link := range links {
		go checkLink(link)
	}
	var input float64
	fmt.Scanf("%f", &input)

}