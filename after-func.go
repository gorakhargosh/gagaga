package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("time elapsed")
		c <- true
	})
	<-c
}
