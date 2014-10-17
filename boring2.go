package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) {
	for i := 0; ; i++ {
		duration := time.Duration(rand.Intn(1e3)) * time.Millisecond
		fmt.Println(msg, i, duration)
		time.Sleep(duration)
	}
}

func main() {
	boring("boring!")
}
