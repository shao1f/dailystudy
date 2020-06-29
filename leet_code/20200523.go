package main

import "fmt"

func FindIndex(a []int) int {
	if len(a) <= 2 {
		return -1
	}
	l := len(a)
	lp, rp := 0, l-1
	lSum, rSum := 0, 0
	for lp < rp {
		if lSum == rSum {
			lSum += a[lp]
			rSum += a[rp]
			lp++
			rp--
		} else if lSum < rSum {
			lSum += a[lp]
			lp++
		} else {
			rSum += a[rp]
			rp--
		}
	}
	if lSum == rSum {
		return lp
	}
	return -1
}

func pivotIndex(a []int) int {
	b := make([]int, 20000)
	b[0] = a[0]
	for i := 1; i < len(a); i++ {
		b[i] = b[i-1] + a[i]
	}
	for i := 1; i < len(a); i++ {
		if b[i-1] == b[len(a)-1]-b[i] {
			return i
		}
	}
	return -1
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	tmp := make([]int, 0, len(nums1)+len(nums2))
	l1, r1, l2, r2 := 0, len(nums1), 0, len(nums2)
	maxLen := max(r1, r2)
	for maxLen > 0 {
		if l1 < r1 && l2 < r2 {

		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// fmt.Println(FindIndex([]int{-1, -1, -1, -1, -1, 0}))
	fmt.Println(pivotIndex([]int{-1, -1, -1, -1, -1, 0}))
}
