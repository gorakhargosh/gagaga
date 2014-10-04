// A channel is a go language construct that provides a mechanism for two
// goroutines to synchronize execution and communicate by passing a value of a
// specified element type. The <- operator specifies the channel direction,
// send or receive. If no direction is given, the channel is bi-directional.
//
// Channels have reference type and are allocated with make (which allocates
// builtin references).
//
// If the channel is unbuffered, the sender blocks until the receiver has
// received the value. If the channel is a buffer, the sender blocks only until
// the value has been copied to the buffer; if the buffer is ful, this means
// waiting until some receiver has retrieved the value. Receivers block until
// there is data to receive.
//
// The close function records that no more values will be sent on a channel.
// After calling close, and after any previously sent values have been
// received, receive operations will return a zero value *without blocking*. A
// multi-valued receive operation additional returns a boolean indicating
// whether the value was delivered by a send operation.
package main

import "fmt"

// We need a type for the sushi to wrap string named values.
type Sushi string

func main() {
	var ch <-chan Sushi = Producer()
	for s := range ch {
		fmt.Println("Consumed", s)
	}
}

func Producer() <-chan Sushi {
	ch := make(chan Sushi)

	go func() {
		ch <- Sushi("海老握り")  // Ebi nigiri
		ch <- Sushi("鮪とろ握り") // Toro nigiri

		// If this channel is not closed, the main goroutine will cause the program
		// to deadlock as the range will never get an indication of ending
		// communication.
		close(ch)
	}()

	return ch
}
