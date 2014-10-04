package main

import "fmt"

func main() {
	// The following program will print "Hello from main goroutine."
	// It /might/ print "Hello from another goroutine." depending upon which
	// of the two goroutines finishes first.
	go fmt.Println("Hello from another goroutine.")
	fmt.Println("Hello from main goroutine.")

	// At this point the program execution stops and all active processes are killed.
}
