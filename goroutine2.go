package main

import (
	"fmt"
	"time"
)

func main() {
	// The following program will print "Hello from main goroutine."
	// It /might/ print "Hello from another goroutine." depending upon which
	// of the two goroutines finishes first.
	go fmt.Println("Hello from another goroutine.")
	fmt.Println("Hello from main goroutine.")

	// Wait a second for the other goroutine to finish.
	// This is not a very nice way to wait for both the goroutines to finish
	// but for the matters of our own perception, this would wait a second for
	// both goroutines to print their messages.
	time.Sleep(time.Second)
}
