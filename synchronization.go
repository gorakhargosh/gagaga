package main

import "fmt"

// Sums an array and sends the resulting value on the channel provided.
func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

// Attempts to sum up an array concurrently.
func main() {
	a := []int{7, 2, 8, -9, 4, 0, 5, 6, 1, 2, 3, 5, 1, -109, 1082, 1983, 1238}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
