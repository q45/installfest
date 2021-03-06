package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/q45/installfest/v2/web"
	"github.com/q45/installfest/v3/model"
	"github.com/q45/installfest/v3/reading"
)

const data = `d2VsY29tZSB0byB0aGUgZ29sZmluZyBtZWV0IHVw`

var (
	message = flag.String("message", "Hello!", "What to say")
	delay   = flag.Duration("delay", 2*time.Second, "how long to wait")
)

//Entry Point for the program
func main() {

	flag.Parse()
	newMessage := *message
	time.Sleep(*delay)

	var g web.Greeting

	r := reading.ReadMyData(newMessage)
	model.MakeItHappen()

	home, _ := time.LoadLocation("America/Utah")

	fmt.Println(home)

	g.Run(r)
}
