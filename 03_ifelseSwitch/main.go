package main

//
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	// 可以没有起始条件和终止条件
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "F"
	case score < 90:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf("Wrong score: %d", score))
	}
	return g

}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func forever() {
	// 死循环
	for {
		fmt.Println("abc")
	}
}

func main() {
	const filename = "abc.txt"
	// if contents, err := ioutil.ReadFile(filename); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%\n", contents)
	// }

	fmt.Println(
		convertToBin(5),
		convertToBin(10),
		convertToBin(12),
	)

	printFile(filename)

	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
		// grade(101),
	)

	a, b := div(13, 3)
	fmt.Println(a, b)

	forever()
}
