package main

import "fmt"

func main() {
	n := factorial(8)
	fmt.Println(n)
}

func factorial(x int) int {

	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}
