package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"
)

const data = `d2VsY29tZSB0byB0aGUgZ29sZmluZyBtZWV0IHVw`

func main() {
	var r io.Reader
	r = strings.NewReader(data)
	r = base64.NewDecoder(base64.StdEncoding, r)
	// r, _ = gzip.NewReader(r)
	b, _ := io.Copy(os.Stdout, r)

	fmt.Println(b)

}
