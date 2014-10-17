package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			duration := time.Duration(rand.Intn(1e3)) * time.Millisecond
			c <- fmt.Sprintf("%s %d %v", msg, i, duration)
			time.Sleep(duration)
		}
	}()
	return c
}

func main() {
	c := boring("joe")
	rand.Seed(time.Now().UTC().UnixNano())
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		// Per communication timeout.
		case <-timeout:
			fmt.Println("Conversation timed out.")
			return
		}
	}
}
