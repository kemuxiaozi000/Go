package main

import (
	"fmt"
	"time"
)

var echo chan string
var receive chan string

// 定义goroutine 1
func Echo() {
	time.Sleep(1 * time.Second)
	echo <- "咖啡色的羊驼"
}

// 定义goroutine 2
func Receive() {
	temp := <-echo // 阻塞等待echo的通道的返回
	receive <- temp
}

func main() {
	echo = make(chan string)
	receive = make(chan string)

	go Echo()
	go Receive()

	getStr := <-receive // 接收goroutine 2的返回

	fmt.Println(getStr)
}
