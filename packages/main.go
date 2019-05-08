package main

import (
	"fmt"

	"github.com/q45/installfest/packages/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("Hello Utah Gophers!"))
	fmt.Println(stringutil.Name)
}
