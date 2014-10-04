package main

import (
	"fmt"
	"time"
)

func BringNews(newsAgency <-chan string, timeout time.Duration) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		select {
		case news := <-newsAgency:
			fmt.Println(news)
		case <-time.After(timeout):
			fmt.Println("Timeout: no news in allotted time.")
		}
		close(wait)
	}()
	return wait
}

func main() {
	toi := make(chan string)
	wait := BringNews(toi, 5*time.Second)
	go func() {
		time.Sleep(10 * time.Second)
		toi <- "No news is good news."
	}()
	<-wait
	fmt.Println("Program khatam thayi gayo.")
}
