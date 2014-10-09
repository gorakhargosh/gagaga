// Synchronization
package main

import (
	"fmt"
	"time"
)

// Publish returns the wait channel and closes the channel after publishing.
func Publish(foo func() string, delay time.Duration) (wait <-chan struct{}) {
	ch := make(chan struct{})
	go func() {
		time.Sleep(delay)
		fmt.Println("BREAKING NEWS:", foo())
		close(ch)
	}()
	return ch
}

func main() {
	wait := Publish(func() string {
		return "Some news."
	}, 5*time.Second)
	fmt.Println("Waiting for the news...")
	<-wait
	fmt.Println("The news is out, time to leave.")
}
