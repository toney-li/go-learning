package main

import (
	"fmt"
)

// Decimal var cannot Printf param as int
func Decimal() {
	fmt.Printf("%d \n", 42)
}

// Binary
func Binary() {
	fmt.Printf("%d - %b \n", 42, 42)
}

func HexaDecimal() {
	fmt.Printf("%d - %b - %x \n", 42, 42, 256)
	fmt.Printf("%d - %b - %#x \n", 42, 42, 256)
	fmt.Printf("%d - %b - %#X \n", 42, 42, 256)
}
func main() {
	Decimal()
	Binary()
	HexaDecimal()
}
