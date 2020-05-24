package main

import "fmt"

// slice 是array 的view ，不是值拷贝，是引用传递

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	// 常用
	arr := [...]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])

	// 引用
	s1 := arr[2:]
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	// cap 可以向后看
	arr1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s2 := arr1[2:6]
	s3 := s2[3:5]
	fmt.Println("cap 可以向后看")
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
	fmt.Printf("s3=%v, len(s3)=%d, cap(s3)=%d\n", s3, len(s3), cap(s3))

	// 当append 超过cap界限会如何
	s4 := append(s2, 10)
	s5 := append(s4, 11)
	s6 := append(s5, 12)
	fmt.Println("当append 超过cap界限会如何")
	fmt.Println("s4,s5,s6=", s4, s5, s6)
	fmt.Println(arr1)

}
