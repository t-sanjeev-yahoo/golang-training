package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	fmt.Println("\n******************************************\nProgram for given Binary number\nIf firstbit is 0 its negative, if 1 is positive\nPrint rest of digits converted to decimal\n*****************************************\n")
	givenNum := "110010"
	actualNum := givenNum[1:]
	firstBit := givenNum[0]
	s := '+'

	if firstBit == '0' {
		s = '-'
	}

	fmt.Printf("givenNum: %v, firstBit: %c, actualNum: %v\n\n", givenNum, firstBit, actualNum)
	
	fmt.Println("Using ParseInt function")
	/*
		func ParseInt(s string, base int, bitSize int) (i int64, err error)
		ParseInt interprets a string s in the given base (0, 2 to 36) and bit size (0 to 64) and returns the corresponding value i.
	*/

	number, _ := strconv.ParseInt(actualNum, 2, 64)

	fmt.Printf("Number: %c%v\n", s, number)

	fmt.Println("\nWithout any function")
	s1 := 0
	n1 := 0
	s2 := 0

	for j, i := 0, len(actualNum)-1; i >= 0; i-- {
		n1 = 0
		if actualNum[j] == '1' {
			n1 = 1
		}
		j++
		s2 = int(math.Pow(2, float64(i))) * n1
		// fmt.Printf("%c,%v,%v,%v\t", actualNum[i], n1, i, s2)
		s1 += s2
	}
	fmt.Printf("Number: %c%v", s, s1)
}
