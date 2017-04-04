package web

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Greeting string

func (g Greeting) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, g)
}

func (g Greeting) Run(r io.Reader) {

	buff := new(bytes.Buffer)
	buff.ReadFrom(r)
	s := buff.String()

	fmt.Println(s)

	err := http.ListenAndServe("localhost:4000", Greeting(s))
	if err != nil {
		log.Fatal(err)
	}
}
