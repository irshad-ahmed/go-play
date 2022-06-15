package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{"http://google.com", "http://amazon.com", "http://golang.org", "http://stackoverflow.com"}

	// creating channel
	c := make(chan string)
	for _, link := range links {
		go checklink(link, c)
	}

	/* 	for {
		go checklink(<-c, c)
	} */

	for l := range c {
		go func(link string) {
			checklink(link, c)
			time.Sleep(5 * time.Second)
		}(l)
	}
}

func checklink(link string, channel chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
		channel <- link
		return
	}
	fmt.Println(link, " is up")
	channel <- link
}
