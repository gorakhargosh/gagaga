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
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		// Per communication timeout.
		case <-time.After(500 * time.Millisecond):
			fmt.Println("You're too slow.")
			return
		}
	}
}
