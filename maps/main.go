package main

import "fmt"

func main() {
	var m map[string]int

	fmt.Println(m)
	if m == nil {
		fmt.Println("m is nil")
	}

	// attempting to add keys to a nil map will result
	// in a runtiem error
	//m["one hundred"] = 100

	var n = make(map[string]int)

	fmt.Println(n)

	if n == nil {
		fmt.Println("n is nil")
	} else {
		fmt.Println("n is not nil")
	}

	//make() function returns an initialed and redy to use map.
	// since it is initialized, you can add new keys

	n["one hundred"] = 100
	fmt.Println(n)

	//initialize with map literal
	m = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	fmt.Println(m)

	//Initialize an empty map
	var q = map[string]int{}
}
