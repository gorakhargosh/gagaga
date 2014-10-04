// Synchronization
package main

import (
	"fmt"
	"time"
)

// Publish returns the wait channel and closes the channel after publishing.
func Publish(text string, delay time.Duration) (wait <-chan struct{}) {
	ch := make(chan struct{})
	go func() {
		time.Sleep(delay)
		fmt.Println("BREAKING NEWS:", text)
		// Causes a deadlock where the main goroutine is expecting to read from
		// this channel and since nothing is being sent on the channel, the receive
		// operation will block forever without allowing the main goroutine to make
		// further progress.
		// close(ch)
	}()
	return ch
}

func main() {
	wait := Publish("Channels let goroutines communicate.", 5*time.Second)
	fmt.Println("Waiting for the news...")
	<-wait
	fmt.Println("The news is out, time to leave.")
}
