package main

import (
	"fmt"
	"sync"
)

type AtomicInt struct {
	mu sync.Mutex
	n  int
}

// Add adds n to the AtomicInt as a single atomic operation.
func (a *AtomicInt) Add(n int) {
	a.mu.Lock()
	a.n += n
	a.mu.Unlock()
}

// Value returns the value of a.
func (a *AtomicInt) Value() int {
	a.mu.Lock()
	n := a.n
	a.mu.Unlock()
	return n
}

// lockItUp demonstrates the use of mutexes to synchronize access to data
// between goroutines.
func lockItUp() {
	wait := make(chan struct{})
	var n AtomicInt
	go func() {
		n.Add(1) // one access.
		close(wait)
	}()
	n.Add(1) // another concurrent access.
	<-wait
	fmt.Println(n.Value())
}

// sharingViaChannelsIsCaring demonstrates the use of channels to synchronize
// access to data between goroutines.
func sharingViaChannelsIsCaring() {
	ch := make(chan int)
	go func() {
		// A local variable is only visible to one goroutine.
		n := 0
		n++
		// The data leaves one goroutine...
		ch <- n
	}()
	// ... and arrives safely in another goroutine.
	n := <-ch
	n++
	fmt.Println(n) // Output guaranteed: 2
}

// race contains a race condition where two goroutines are competing for access
// to a shared variable.
func race() {
	wait := make(chan struct{})
	n := 0
	go func() {
		n++ // one access: read, increment, write
		close(wait)
	}()
	// Another conflicting access.
	n++
	<-wait
	fmt.Println(n) // May output: 2, 1, etc.
}

func main() {
	race()
	sharingViaChannelsIsCaring()
}
