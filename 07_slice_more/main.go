package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("s=%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {

	// 创建slice
	// type 1
	var s []int
	for i := 0; i < 100; i++ {
		// printSlice(s)
		s = append(s, 2*i+1)
	}
	printSlice(s)

	// type 2
	s1 := []int{2, 4, 6, 8, 10}
	printSlice(s1)

	// type 3
	s2 := make([]int, 10)
	printSlice(s2)

	s3 := make([]int, 10, 32)
	printSlice(s3)
	//copy
	fmt.Println("Copying slice")
	copy(s2, s1)
	printSlice(s2)

	//delete
	fmt.Println("Delete element from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	//pop from front
	fmt.Println("pop from front")
	// front := s2[0]
	s2 = s2[1:]
	printSlice(s2)

	//pop from back
	fmt.Println("pop from back")
	// tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	printSlice(s2)

	//pop
}
