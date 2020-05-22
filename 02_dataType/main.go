package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func euler() {
	// c := 3 + 4i
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	// cmplx.Abs(c)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	euler()
	triangle()
}
