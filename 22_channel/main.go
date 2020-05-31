package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("worker %d received %c\n", id, n)
	}
	// 第二种channel close 判断方法
	// for n:=range c{
	// 	fmt.Printf("worker %d received %c\n", id, n)
	// }
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func channelDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
		// n := <-channels[i]
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd' //all goroutines are asleep - deadlock!
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	close(c) // 收空串
	time.Sleep(time.Millisecond)
}

func main() {
	// channelDemo()
	// bufferedChannel()
	channelClose()
}
