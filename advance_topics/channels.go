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
	//Declaring Channel
	channel := make(chan string)
	for _, link := range links {
		go checkLink(link, channel)
	}
	for i := 0; i < len(links); i++ {
		fmt.Println(<-channel)
	}

}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link)
	if err != nil {
		//fmt.Println(link, "Might be down with ", err)
		channel <- link + "Might be down with "
		return
	}
	//fmt.Println(link, "is up!")
	channel <- link + "is up!"

}
