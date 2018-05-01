package main

import (
	"fmt"
	"time"

	"github.com/installfest/model"
	"github.com/installfest/reading"
	"github.com/installfest/web"
)

const data = `d2VsY29tZSB0byB0aGUgZ29sZmluZyBtZWV0IHVw`

//entry point for this app
func main() {

	var g web.Greeting

	r := reading.ReadMyData(data)
	model.MakeItHappen()

	home, _ := time.LoadLocation("America/Utah")

	fmt.Println(home)

	g.Run(r)
}
