package main

import "fmt"

func subarraysDivByK(A []int, K int) int {
	B := make([]int, len(A), len(A))
	B[0] = A[0]
	l := len(A)
	for i := 1; i < l; i++ {
		B[i] = B[i-1] + A[i]
	}
	for j := 0; j < l; j++ {
		// fmt.Println(B[j] - B[j-1])
		for i := j + 1; i < l; i++ {
			fmt.Println(B[i] - B[j])
		}
	}
	fmt.Println(B)
	return 0
}

func subarraysDivByKOne(A []int, K int) int {
	l := len(A)
	if l <= 0 {
		return 0
	}
	res := 0
	for j := 1; j <= l; j++ {
		for i := 0; i < j; i++ {
			if SumArr(A[i:j], K) {
				res++
			}
		}
	}
	return res
}

func SumArr(SubA []int, K int) bool {
	sum := 0
	for _, v := range SubA {
		sum += v
	}
	if sum%K == 0 {
		fmt.Println(SubA)
		return true
	}
	return false
}

func main() {
	// fmt.Println(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5))
	fmt.Println(-4 % 5)
}
