package main

import "fmt"

func main() {
	fmt.Println(Factorial(8))
}

func Factorial(num int) int {

	var sum []int
	// var factorial int

	for i := num; i > 0; i-- {
		fmt.Println(i)
		sum = append(sum, i)
	}

	fmt.Println(sum)

	// for _, f := range sum {

	// }

	return 0

}
