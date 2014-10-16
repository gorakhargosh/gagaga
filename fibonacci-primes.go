package main

import "fmt"

// isPrime determines whether a specified integer is a prime number.
// This is a very naive implementation of the prime determinator function.
func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// A Fibonacci series generator using goroutines and channels for concurrent
// series generation. Python's "generator" pattern impl.
func fibonacci(n int) chan int {
	c := make(chan int)
	go func() {
		a, b := 0, 1
		for i := 0; i < n; i++ {
			a, b = b, a+b
			c <- a
		}
		// Prevents a deadlock and sends EOC to receiver.
		close(c)
	}()
	return c
}

// filterPrimes filters values transmitted between input and output channels
// depending upon whether the values are prime integers.
func filterPrimes(cin chan int) chan int {
	cout := make(chan int)
	go func() {
		for x := range cin {
			if isPrime(x) {
				cout <- x
			}
		}
		// Close channel to prevent deadlock and send EOC to receiver.
		close(cout)
	}()
	return cout
}

func main() {
	n := 10
	fmt.Printf("Fibonacci primes until %v fibonacci numbers dumped.\n", n)
	fmt.Println(isPrime(2), isPrime(201), isPrime(131), isPrime(4))
	for x := range filterPrimes(fibonacci(n)) {
		fmt.Println(x)
	}
	fmt.Println("EOP")
}
