package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		duration := time.Duration(rand.Intn(1e3)) * time.Millisecond
		c <- fmt.Sprintf("%s %d %v", msg, i, duration)
		time.Sleep(duration)
	}
}

func main() {
	// Running the boring function as a goroutine.
	c := make(chan string)
	go boring("boring!", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring. I'm leaving.")
}
