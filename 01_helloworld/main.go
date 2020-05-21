package main

import "fmt"

//函数的外面只能放标识符的声明
// fmt.Println("Hello") //非法

var (
	name string
	age  int
	isOK bool
)

const (
	n1 = 100
	n2
	n3
)
const (
	a1 = iota //0
	a2 = 100  //100
	a3 = iota //2
	a4        //4
)

// 定义数量级、
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	name = "haha"
	age = 17
	isOK = true
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)
	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)
	// 变量声明必须使用、
}
