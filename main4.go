package main

import (
	"fmt"
)

func main4() {
	s := "Hello, 世界"
	for i, r := range s {
		fmt.Printf("Index: %d, Rune: %c, %T\n", i, r, r)
	}

	fmt.Println()
	s = "Hello, GO"
	for i, r := range s {
		fmt.Printf("Index: %d, Rune: %c, %T\n", i, r, r)
	}

	fmt.Println()
	var s1 string = "Hello, 世界"
	for i, r := range s1 {
		fmt.Printf("Index: %d, Rune: %v, %T\n", i, r, r)
	}

	fmt.Println()
	var s2 string = "Hello, GO"
	for i, r := range s2 {
		fmt.Printf("Index: %d, Rune: %v, %T\n", i, r, r)
	}

	var b1 byte = 'g'
	fmt.Printf("b1: %v, %T\n", b1, b1)

	var b2 = []byte{'g', 'g', 'g'}
	fmt.Printf("b2: %v, %T\n", b2, b2)

	fmt.Printf("b1: %v, %T\n", b1, b1)
	c := string(b2)
	fmt.Printf("b1: %v, %T\n", c, c)
}
