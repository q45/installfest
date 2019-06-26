package main

import "fmt"

//lowercase non-exported
type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 36})

	s := person{name: "Janice", age: 44}

	s.age = 99
	fmt.Println(s)

	s.getName()
}

func (p *person) getName() {
	fmt.Println(p.name)
}
