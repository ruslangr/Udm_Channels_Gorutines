package main

import (
	"fmt"
	"net/http"
	"time"
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

	// для блокировки завершения добавляем цикл с ожиданием вывода из
	//	for i := 0; i < len(links); i++ {
	//		fmt.Println(<-c)
	//	}

	for l := range c {
		//	time.Sleep(5 * time.Second) - если добавить сюда, то будет паузится main gourutine, соотвественно ответу от вспомогательных gorutine некуда будет приходить
		//go checkLink(l, c)
		go func() {
			time.Sleep(5 * time.Second)
			go checkLink(l, c)
		}() // - скобки нужны для запуска анонимной функции
	}

}

func checkLink(link string, c chan string) {
	//	time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "- not available", err)
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	//c <- "Yep its up" //делается для блокирования и препятствования преждевременного завершения main
	c <- link
}
