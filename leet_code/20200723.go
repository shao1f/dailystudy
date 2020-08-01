package main

import "fmt"

func removeDuplicates(nums []int) int {
	slow := 0
	for i := 0; i < len(nums); i++ {
		if nums[slow] != nums[i] {
			fmt.Println(slow)
			nums[slow] = nums[i]
			slow++
		}
	}
	return slow
}

func main() {
	removeDuplicates([]int{1, 1, 2})
}
