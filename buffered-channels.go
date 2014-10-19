package main

import "fmt"

func main() {
	c := make(chan int, 2)

	// These send operations won't block.
	c <- 1
	c <- 2

	// And these lines will be executed.
	fmt.Println(<-c)
	fmt.Println(<-c)

	c <- 1
	c <- 2
	c <- 3 // This will cause a deadlock because the buffer is full and values
	// have not been read yet.

	fmt.Println(<-c)
	fmt.Println(<-c)
}
