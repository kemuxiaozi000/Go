package main

import (
	"fmt"
	"sync"
)

// func doWork(id int, c chan int, wg *sync.WaitGroup) {
func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("worker %d received %c\n", id, n)
		// done <- true
		w.done()
	}
}

type worker struct {
	in chan int
	// wg *sync.WaitGroup
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	// c := make(chan int)
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	// 这里的wg已经是指针了，不需要&wg
	go doWork(id, w)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup

	var workers [10]worker

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	wg.Wait()
}

func main() {
	chanDemo()
}
