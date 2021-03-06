package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.golang.org",
		"https://www.twitter.com",
	}
	//Declaring Channel
	fmt.Println("Starting")
	channel := make(chan string)
	for _, link := range links {
		go checkLink(link, channel)
	}

	for l := range channel {
		//to add a pause
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, channel)
		}(l)

	}

}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "Might be down with ", err)
		channel <- link
		return
	}
	fmt.Println(link, "is up!")
	channel <- link

}
