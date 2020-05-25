package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"course1": "golang1",
		"course2": "golang2",
		"course3": "golang3",
	}
	m2 := make(map[string]int)
	var m3 map[string]int
	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Get value")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)

	caurseName, ok := m["caurse"]
	fmt.Println(caurseName, ok) // false
}
