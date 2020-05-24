package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args"+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	fmt.Println(numbers)
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	fmt.Println(apply(pow, 4, 2))
	fmt.Println(sum(1, 4, 2))

}
