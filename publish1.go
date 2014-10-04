package main

import (
	"fmt"
	"time"
)

// Publish prints text to stdout after the given time has elapsed. It doesn't
// block but returns right away.
func Publish(text string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println("BREAKING NEWS:", text)
	}()
}

func main() {
	Publish("A goroutine starts a new thread of execution.", 5*time.Second)
	fmt.Println("Let's hope the news will be published before I leave.")

	// Change this duration to a value less than the waiting time for the main
	// thread and you will not see the breaking news!
	durationInSeconds := 10 * time.Second
	// Wait for the news to be published.
	time.Sleep(durationInSeconds)
	fmt.Println("Ten seconds later: I'm leaving now.")

	// In general it's not possible to arrange for threads to wait for each other
	// by sleeping. In order to synchronize the communication and the waiting
	// between goroutines, we need channels.
}
