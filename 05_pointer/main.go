package main

import "fmt"

func swap(a, b int) {
	b, a = a, b
}
func swap1(a, b *int) {
	*b, *a = *a, *b
}

// 推荐直接返回
func swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 3, 4
	swap(a, b)
	fmt.Println(a, b)

	c, d := 3, 4
	swap1(&c, &d)
	fmt.Println(c, d)

	e, f := 3, 4
	fmt.Println(swap2(e, f))
}
