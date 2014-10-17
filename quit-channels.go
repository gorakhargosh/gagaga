package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, quit chan bool) chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			duration := time.Duration(rand.Intn(1e3)) * time.Millisecond
			select {
			case c <- fmt.Sprintf("%s %d %v", msg, i, duration):
				// do nothing
			case <-quit:
				// Exit function.
				// do some clean up here.
				quit <- true
				return
			}
			time.Sleep(duration)
		}
	}()
	return c
}

func main() {
	quit := make(chan bool)

	c := boring("Joe", quit)

	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
	<-quit
}
