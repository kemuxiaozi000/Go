package main

import "fmt"

func main() {
	// deadlock1()
	// deadlock2()
	useSelect()
}

func deadlock1() {
	ch := make(chan int)
	<-ch // 阻塞main goroutine, 通道被锁
}
func deadlock2() {
	cha, chb := make(chan int), make(chan int)

	go func() {
		cha <- 1 // cha通道的数据没有被其他goroutine读取走，堵塞当前goroutine
		chb <- 0
	}()

	<-chb // chb 等待数据的写
}

func useSelect() {
	ch := make(chan int, 1)

	ch <- 1
	select {
	case <-ch:
		fmt.Println("咖啡色的羊驼")
	case <-ch:
		fmt.Println("黄色的羊驼")
	}
}
