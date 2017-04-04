package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/installfest/v3/model"
	"github.com/installfest/v3/reading"
	"github.com/installfest/v3/web"
)

const data = `d2VsY29tZSB0byB0aGUgZ29sZmluZyBtZWV0IHVw`

var (
	message = flag.String("message", "Hello!", "What to say")
	delay   = flag.Duration("delay", 2*time.Second, "how long to wait")
)

func main() {

	flag.Parse()
	fmt.Println(*message)
	time.Sleep(*delay)

	r := reading.ReadMyData(data)
	model.MakeItHappen()

	home, _ := time.LoadLocation("America/Utah")

	fmt.Println(home)

	var g web.Greeting

	g.Run(r)
}
