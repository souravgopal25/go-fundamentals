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
		go checkLink1(link)
	}
	var input float64
	fmt.Scanf("%f", &input)

}

func checkLink1(link string) {
	_, err := http.Get(link)
	if err != nil {
		//fmt.Println(link, "Might be down with ", err)
		fmt.Println(link, " Might be down with ", err)
		return
	}
	fmt.Println(link, "is up!")
	//channel <- link + "is up!"

}
