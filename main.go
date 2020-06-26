package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
		//	fmt.Println(<-c) // main не завершите пока не получит из канала сообщение
	}
	//fmt.Println(<-c) - если сделать не в цикле, то программа завершится после получения 1го результата
	//fmt.Println(<-c) - или завершится после получения 2го

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "- not available", err)
		c <- "May be a problem"
		return
	}
	fmt.Println(link, "is up!")
	c <- "Yep its up" //делается для блокирования и препятствования преждевременного завершения main
}
