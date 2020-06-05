package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// 传统同步机制

func main() {
	request, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
	request = request.Header.Add()
	resp, err := http.DefaultClient.Do(request)
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}
