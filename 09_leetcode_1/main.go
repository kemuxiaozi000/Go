package main

import "fmt"

func lengtOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {

		if lastI, ok := lastOccurred[ch]; ok && lastOccurred[ch] >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

// 有bug
func draft(s string) int {
	maxLen, start := 0, 0
	m := make(map[rune]int)
	for k, v := range s {
		// 看有没有 m[v] err = false 代表没有
		currentIdx, err := m[v]
		m[v] = k
		if err {
			fmt.Println("k-currentIdx=", k-currentIdx)
			if k-currentIdx > maxLen && k-start > maxLen {
				maxLen = k - currentIdx + 1
			}
			start = k
		} else {
			if k-start > maxLen {
				maxLen++
			}

		}
		fmt.Println("start=", start)
		fmt.Println("m=", m)
		fmt.Println("maxLen=", maxLen)
		fmt.Println("----------------")
	}
	return maxLen
}

func main() {
	// s := "abcabcbcbcbca"
	// fmt.Println(lengtOfNonRepeatingSubStr("abcabcbb"))
	// fmt.Println(lengtOfNonRepeatingSubStr("bbbbb"))
	// fmt.Println(lengtOfNonRepeatingSubStr("pwwkew"))

	// fmt.Println(maxIdx, maxLength)
	// fmt.Println(draft("abc"))
	// fmt.Println(draft("abcabcbb"))
	// fmt.Println(draft("bbbbb"))
	fmt.Println(draft("pwwkew"))

}
