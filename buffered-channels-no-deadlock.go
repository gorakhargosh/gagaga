package main

import "fmt"

func main() {
	c := make(chan int, 2)

	c <- 1
	c <- 2
	// The following will not deadlock because the chan send operation is
	// scheduled on a separate thread.
	c3 := func() { c <- 3 }
	go c3()
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
