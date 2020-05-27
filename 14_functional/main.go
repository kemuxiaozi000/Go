package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occured")
	fmt.Println(4)
}

func writeFile(fliename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	// f := fib.Fibonacci()

	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, 1)
	}
}

func main() {
	tryDefer()
	writeFile("fib.txt")
}
