package main

import (
	"fmt"
	"time"

	"./mock"
	"./real"
)

// Retriver xxx
type Retriver interface {
	Get(url string) string
}

// Poster xxx
type Poster interface {
	Post(url string, form map[string]string) string
}

// RetriverPoster 组合接口
type RetriverPoster interface {
	Retriver
	Poster
}

func post(poster Poster) {
	poster.Post("http://www.baidu.com",
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

func download(r Retriver) string {
	return r.Get("http://www.baidu.com")
}

func session(s RetriverPoster) string {
	s.Post("http://www.baidu.com", map[string]string{
		"contents": "哈哈哈",
	})
	return s.Get("http://www.baidu.com")
}

func inspect(r Retriver) {
	switch v := r.(type) {
	case *mock.Retriver:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}

}

func main() {
	var r Retriver
	retr := &mock.Retriver{"this is a mock"}
	// fmt.Printf("%T %v\n", r, r)
	r = retr
	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	// fmt.Printf("%T %v\n", r, r)
	// fmt.Println(download(r))

	// type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)
	fmt.Println("组合interface")
	fmt.Println(session(retr))
}
