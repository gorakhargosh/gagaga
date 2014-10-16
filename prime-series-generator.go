package main

import "fmt"

// isPrime determines whether a number is a prime number.  Incredibly slow way
// to determine whether a number is prime with naive impl.
func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// primes generates a series of prime numbers sending them on a channel that it
// returns.
func primes(n int) chan int {
	c := make(chan int)
	go func() {
		for i := 2; n > 0; i++ {
			if isPrime(i) {
				n--
				c <- i
			}
		}
		close(c)
	}()
	return c
}

// Entry point.
func main() {
	n := 10
	fmt.Printf("Generates a series of %v prime numbers concurrently.\n", n)

	for x := range primes(n) {
		fmt.Println(x)
	}
	fmt.Println("EOP")
}
