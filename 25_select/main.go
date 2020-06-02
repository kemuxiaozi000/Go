package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
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
	n := 0
	hasValue := false
	for {
		var activeWorker chan int // nil channel
		select {
		case <-c1:
			w <- n
			// fmt.Println("Receive from c1", n)
			hasValue = true
		case n := <-c2:
			w <- n
			hasValue = true
		// fmt.Println("Receive from c2", n)
		// default:
		// 	fmt.Println("no value ")
		case w <= n:
			if hasValue {

			}

		}

	}

}
