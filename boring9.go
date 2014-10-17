package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str    string
	resume chan bool
}

// Generator pattern.
func boring(msg string) chan Message {
	c := make(chan Message)
	resume := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			duration := time.Duration(rand.Intn(1e3)) * time.Millisecond
			c <- Message{fmt.Sprintf("%s %d %v", msg, i, duration), resume}
			time.Sleep(duration)
			<-resume
		}
	}()
	return c
}

func fanIn(a, b <-chan Message) <-chan Message {
	c := make(chan Message)
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
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.resume <- true
		msg2.resume <- true
	}
	fmt.Println("You're boring. I'm leaving.")
}
