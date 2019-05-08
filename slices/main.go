package main

import "fmt"

//SLICES & Arrays
/*
	Array has a fixed size.

	Slice is a dynamically-sized, flexible view
	into the elements of an array
	Slices are much more commong than arrays

	A slice contians 3 things
	 A pointer to an underlying array.
	 The length of the segment of the array that the slice contains
	 the cpacity (maximum size up to which the segment can grow)
*/

func main() {
	primes := [6]int{2, 3, 4, 5, 11, 13}

	// a[low : high]
	// this selects a half-open range which includes the first element,
	// but excludes the last one
	var s = primes[1:4]

	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	fmt.Println(s)

	//Create an array of size 5, and return a slice that references it
	r := make([]int, 5)
	fmt.Printf("s = %v, len = %d, cap = %d\n", r, len(r), cap(r))

	//Zero Value Slice
	var z []int
	fmt.Printf("%z = %v, len = %d, cap = %d\n", z, len(z), cap(z))

	if z == nil {
		fmt.Println("z is nil")
	}

	var animals []string
	animals = append(animals, "Cat", "Dog", "Lion", "Tiger")

	fmt.Printf("b = %v, len = %d, cap = %d\n", animals, len(animals), cap(animals))

	for i, v := range animals {
		fmt.Printf("index = %d\n value = %s\n", i, v)
	}

}
