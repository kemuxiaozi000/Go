package main

import (
	"fmt"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			// time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			time.Sleep(5 * time.Second)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	// 非阻塞式 select + default
	var w chan int
	var values []int
	n := 0
	hasValue := false
	for {
		var activeWorker chan int // nil channel
		var activeValue int
		if hasValue {
			activeWorker = w
			activeValue = values[0]
		}
		select {
		case n = <-c1:
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
		// nil 的情况下不会执行
		case activeWorker <- activeValue:
			//把values第一个值拿走
			values = values[1:]
		}

	}

}
