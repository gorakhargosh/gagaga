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

func main() {
	// Running the boring function as a goroutine.
	c := boring("boring!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring. I'm leaving.")
}
