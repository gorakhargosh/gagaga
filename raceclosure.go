package main

import (
	"fmt"
	"sync"
)

func main() {
	race()
	correct()
	alsoCorrect()
}

func race() {
	var w sync.WaitGroup
	w.Add(5)
	for i := 0; i < 5; i++ {
		// This variable i is shared by 6 (six) goroutines and is therefore prone
		// to race conditions.
		go func() {
			fmt.Println(i)
			w.Done()
		}()
	}
	w.Wait()
	fmt.Println()
}

// Use a local variable and pass the number as a parameter so every goroutine
// gets its own copy of the value.
func correct() {
	var w sync.WaitGroup
	w.Add(5)
	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Println(n)
			w.Done()
		}(i)
	}
	w.Wait()
	fmt.Println()
}

// Also correct if you want to keep using a closure, but creates a copy of the
// variable in another way per goroutine. Same difference. Passing by argument
// is just more elegant.
func alsoCorrect() {
	var w sync.WaitGroup
	w.Add(5)
	for i := 0; i < 5; i++ {
		// Make a copy of the variable outside the closure and use that.
		n := i
		go func() {
			fmt.Println(n)
			w.Done()
		}()
	}
	w.Wait()
	fmt.Println()
}
