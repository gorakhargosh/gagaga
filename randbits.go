package main

import "fmt"

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
	for k := range RandomBits() {
		fmt.Println(k)
	}
}
