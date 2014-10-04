package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	// Values cannot be mutated.
	zeroval(i)
	fmt.Println("zeroval:", i)

	// Mutation via pointers.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// To see what a pointer value looks like in memory.
	fmt.Println("pointer:", &i)
}
