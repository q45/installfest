package model

import (
	"encoding/json"
	"fmt"
	"strings"
)

const blob = `[
		{"Title": "0redev", "URL":"http://ordedev.org"},
		{"Title": "Strange Loop", "URL": "http://thestrangeloop.com"}
	]`

type Item struct {
	Title string
	URL   string
}

func MakeItHappen() []*Item {

	//Type Safety
	var items []*Item
	json.NewDecoder(strings.NewReader(blob)).Decode(&items)
	for _, item := range items {
		fmt.Printf("Title: %v URL: %v\n", item.Title, item.URL)
	}

	return items
}
