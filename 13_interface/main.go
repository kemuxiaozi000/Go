package main

import (
	"fmt"

	"./real"
)

// Retriver xxx
type Retriver interface {
	Get(url string) string
}

func download(r Retriver) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriver
	// r = mock.Retriver{"this is a mock"}
	r = real.Retriever{}
	fmt.Println(download(r))
}
