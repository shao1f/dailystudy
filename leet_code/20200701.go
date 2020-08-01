package main

import "fmt"

func FindLength(a, b []int) int {
	lenA := len(a)
	lenB := len(b)
	if lenA == 0 || lenB == 0 {
		return 0
	}
	stak := make([]int, 0, 0)
	// stak = append(stak, a[0])
	resLen := 0
	i, j := 0, 0
	for i < lenA && j < lenB {
		if a[i] == b[j] {
			stak = append(stak, a[i])
			i++
			j++
		} else {
			curLen := len(stak)
			if curLen > resLen {
				resLen = curLen
			}
			stak = stak[:0]
			i++
		}
	}
	return len(stak)
}

func main() {
	a := []int{0, 0, 0, 0, 1}
	b := []int{1, 0, 0, 0, 0}
	fmt.Println(FindLength(a, b))
}
