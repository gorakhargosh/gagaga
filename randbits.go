package main

import "fmt"

// RandomBits generates random bits and returns a channel that can be used
// to read the random bits from.
func RandomBits() <-chan int {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- 0: // note: no statement
			case ch <- 1:
			}
		}
	}()
	return ch
}

func main() {
	// Be aware that this function will never end because both communications are
	// ready at all times and the select will pseudorandomly select either one of
	// them. Therefore, there is no termination condition for range to cause the
	// loop to exit.
	for k := range RandomBits() {
		fmt.Println(k)
	}
}
