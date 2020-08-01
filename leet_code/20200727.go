package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) == 0 {
		return false
	}
	resMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if resMap[nums[i]] == 0 {
			resMap[nums[i]] = i + 1
		} else {
			if i-resMap[nums[i]] < k {
				return true
			} else {
				resMap[nums[i]] = i + 1
			}
		}
	}
	return false
}

func quickSort(nums *[]int, left, right int) {
	if left > right {
		return
	}
	s := (*nums)[left]
	i, j := left, right
	for {
		for i < j && s <= (*nums)[j] {
			j--
		}
		for i < j && s >= (*nums)[i] {
			i++
		}
		if i >= j {
			break
		}
		(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
	}
	(*nums)[left], (*nums)[i] = (*nums)[i], (*nums)[left]
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}

func main() {
	// containsNearbyDuplicate([]int{1, 0, 1, 1}, 1)
	nums := []int{3, 5, 1, 2, 6}
	quickSort(&nums, 0, len(nums)-1)
	fmt.Println(nums)
}
