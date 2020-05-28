package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	err = errors.New("this is a custom error")
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, "1")
	}
}

func main() {
	// tryDefer()
	writeFile("fib.txt")
}
