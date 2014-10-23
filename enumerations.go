package main

import (
	"fmt"
	"math"
)

const (
	FlagZero = 0
	FlagOne  = 4
	FlagTwo
)

const (
	Cyan = iota
	Yellow
	Magenta = 40
	Blue    = iota
	Green
	Red
	Black
	Violet
)

const (
	First = 2
	Second
)

const (
	StatusOn = 1 << iota
	StatusOff
	StatusInactive
	StatusFound
	StatusNotFound
	StatusAccessDenied
)

func main() {
	fmt.Println(Cyan, Magenta, Yellow, Blue, Green, Red, Black, Violet)
	fmt.Println(First, Second)
	fmt.Println(
		StatusOn, StatusOff, StatusInactive, StatusFound,
		StatusNotFound, StatusAccessDenied)
	fmt.Println(FlagZero, FlagOne, FlagTwo)
	fmt.Println(math.Pi)
}
