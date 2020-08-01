package main

import (
	"fmt"
	"sort"
)

// 求二进制1的个数
func binaryCount(a int) int {
	count := 0
	for a > 0 {
		if a&1 == 1 {
			count++
		}
		a >>= 1
	}
	return count
}

func A(a []int) {
	sort.Slice(a, func(i, j int) bool {

	})
}

func main() {
	fmt.Println(binaryCount(3))
	b := func(a int) int {
		count := 0
		for a > 0 {
			if a&1 == 1 {
				count++
			}
			a >>= 1
		}
		return count
	}
	fmt.Println(b(5))

}
