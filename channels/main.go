package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com", "http://facebook.com", "http://amazon.com",
	}

	c := make(chan string)
	for _, link := range links {
		go checkLinks(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLinks(link, c)
		}(l)
	}

}

func checkLinks(link string, c chan string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println("Error:", err)
		c <- link
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode, link, "is active")
	c <- link
}
