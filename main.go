package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/installfest/model"
	"github.com/installfest/reading"
	"github.com/installfest/web"
)

const data = `d2VsY29tZSB0byB0aGUgZ29sZmluZyBtZWV0IHVw`

var (
	message = flag.String("message", "Hello!", "What to say")
	delay   = flag.Duration("delay", 2*time.Second, "how long to wait")
)

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
