package main

import (
	"fmt"
	"time"
)

func doWork(id int, c chan int, done chan bool) {
	// 第一种channel close 判断方法
	// for {
	// 	n, ok := <-c
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Printf("worker %d received %c\n", id, n)
	// }

	// 第二种channel close 判断方法
	for n := range c {
		fmt.Printf("worker %d received %c\n", id, n)
		// done <- true 会导致如下问题
		// fatal error: all goroutines are asleep - deadlock!
		// 循环等待
		// 或者两个任务分开等
		go func() { done <- true }()
	}
}

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker {
	// c := make(chan int)
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w.in, w.done)
	return w
}

func channelDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
		// <-workers[i].done
		// n := <-channels[i]
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
		// <-workers[i].done
	}
	// time.Sleep(time.Millisecond)
	// wait for all of them
	// 两个任务一起等
	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}

}

func bufferedChannel() {
	c := make(chan int, 3)
	// go doWork(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd' //all goroutines are asleep - deadlock!
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	// go doWork(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	close(c) // 收空串
	time.Sleep(time.Millisecond)
}

func main() {
	channelDemo()
	// bufferedChannel()
	// channelClose()
}
