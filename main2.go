package main // main package

import "fmt" // import fmt package

func main2() { // main function
	fmt.Printf("Hello World..\nWelcome to GO Programming...\n")

	var c1 map[string]string

	var c2 = map[string]string{

		"USA":    "+1",
		"India":  "+91",
		"China":  "+86",
		"Brazil": "+55",
		"UK":     "+44",
	}
	var c3 = map[string]int{
		"USA":    +1,
		"India":  +91,
		"China":  +86,
		"Brazil": +55,
		"UK":     +44,
	}
	c4 := map[string]float32{
		"USA":    1.11,
		"India":  91.2111,
		"China":  0.86,
		"Brazil": 5.555555555555,
		"UK":     44,
	}

	c5 := map[int]float32{
		1:  1.11,
		4:  91.2111,
		5:  0.86,
		6:  5.555555555555,
		12: 44,
	}
	fmt.Println(c1)
	// for k, v := range c1 {
	// 	fmt.Printf("k:%v, v:%v\n", k, v)
	// }
	fmt.Println(c2)
	for k, v := range c2 {
		fmt.Printf("k:%v, v:%v\n", k, v)
	}
	fmt.Println(c3)
	for k, v := range c3 {
		fmt.Printf("k:%v, v:%v\n", k, v)
	}
	fmt.Println(c4)
	for k, v := range c4 {
		fmt.Printf("k:%v, v:%v\n", k, v)
	}
	fmt.Println(c5)
	for k, v := range c5 {
		fmt.Printf("k:%v, v:%v\n", k, v)
	}
}
