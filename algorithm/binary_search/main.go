package main

import "fmt"

func binarySearch(nums []int, num int) bool {
	l, r := 0, len(nums)
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] > num {
			r = mid - 1
		} else if nums[mid] < num {
			l = mid + 1
		} else {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	find := binarySearch(nums, 5)
	fmt.Println(find)
}
