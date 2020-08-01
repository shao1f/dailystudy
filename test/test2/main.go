package main

import "fmt"

func reverseArr(arr [][]int) [][]int {
	res := make([][]int, 0, len(arr))
	if len(arr) == 0 {
		return res
	}
	for i := 0; i < len(arr[0]); i++ {
		tmp := make([]int, 0, len(arr[0]))
		for j := len(arr) - 1; j >= 0; j-- {
			tmp = append(tmp, arr[j][i])
		}
		res = append(res, tmp)
	}
	return res
}

func reverse(arr [][]int) {
	for i := 0; i < len(arr[0]); i++ {
		for j := len(arr) - 1; j >= 0; j-- {

		}
	}
}

func main() {
	arr1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(reverseArr(arr1))
}
