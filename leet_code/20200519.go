package main

import "fmt"

func validPalindrome(s string) bool {
	l := len(s)
	if l <= 0 {
		return false
	}
	low := 0
	high := l - 1
	for low < high {
		// ii := string(s[i])
		// jj := string(s[j])
		// _ = ii
		// _ = jj
		if s[low] == s[high] {
			low++
			high--
		} else {
			var flag1, flag2 = true, true
			i1, j1 := low, high-1
			for i1 < j1 {
				if s[i1] != s[j1] {
					flag1 = false
				}
				i1++
				j1--
			}
			i2, j2 := low+1, high
			for i2 < j2 {
				if s[i2] != s[j2] {
					flag2 = false
				}
				i2++
				j2--
			}
			return flag1 || flag2
		}
	}
	return true
}

func main() {
	// i := 0
	// j := 10
	// for i <= j {
	// 	i++
	// 	j--
	// 	fmt.Println(i, ":", j)
	// }
	fmt.Println(validPalindrome("ebcbbececabbacecbbcbe"))
}
