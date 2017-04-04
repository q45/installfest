package reading

import (
	"encoding/base64"
	"io"
	"strings"
)

func ReadMyData(data string) io.Reader {
	var r io.Reader
	r = strings.NewReader(data)
	r = base64.NewDecoder(base64.StdEncoding, r)
	// r, _ = gzip.NewReader(r)
	// b, _ := io.Copy(os.Stdout, r)

	return r
}
