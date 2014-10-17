package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generator pattern.
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

func fanIn(a, b <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-a
		}
	}()
	go func() {
		for {
			c <- <-b
		}
	}()
	return c
}

func main() {
	// Running the boring function as a goroutine.
	c := fanIn(boring("joe"), boring("ann"))
	for i := 0; i < 15; i++ {
		// Joe and ann services are not synchronized but no multiplexing.
		// Joe blocks reading from ann even if ann is ready. We need a fan-in
		// multiplexer.
		fmt.Println(<-c)
	}
	fmt.Println("You're boring. I'm leaving.")
}
