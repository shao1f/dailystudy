package main

import "fmt"

type TestInt int

func (t TestInt) Valid() bool {
	if t == 0 || t%2 == 0 {
		return true
	}
	return false
}

func findTheLongest(s string) int {
	curLen := 0
	tmpMap := make(map[string]TestInt)
	for i := 0; i < len(s); i++ {
		tmpMap[string(s[i])]++
		if tmpMap["a"].Valid() && tmpMap["e"].Valid() && tmpMap["i"].Valid() && tmpMap["o"].Valid() && tmpMap["u"].Valid() {
			curLen = i + 1
		}
	}
	return curLen
}

func findTheLongestSubstring(s string) int {
	maxLen := 0
	for i := 0; i < len(s); i++ {
		l := findTheLongest(s[i:])
		maxLen = Max(l, maxLen)
	}
	return maxLen
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test1(s []int, l, r int) int {
	sumL := 0
	sumR := 0
	for i, v := range s {
		if i < l {
			sumL += v
		}
		if i <= r {
			sumR += v
		}
	}
	return sumR - sumL
}

func PreSumOne() {
	a := make([]int, 1000, 1000)
	b := make([]int, 1000, 1000)
	var index int
	fmt.Print("please enter index:")
	fmt.Scanln(&index)
	for i := 1; i <= index; i++ {
		fmt.Print("please enter number:")
		fmt.Scanln(&a[i])
		b[i] = b[i-1] + a[i]
	}
	var t int
	fmt.Print("please enter t:")
	fmt.Scanln(&t)
	for t > 0 {
		var l, r int
		fmt.Print("please enter l:")
		fmt.Scanln(&l)
		fmt.Print("please enter r:")
		fmt.Scanln(&r)
		if l > len(a) || r > len(b) {
			continue
		}
		res := b[r] - b[l-1]
		fmt.Println("res:", res)
		t--
	}
}

// func PreSumTwo() {
// 	a := make([]int, 1000, 1000)
// 	b := make([]int, 1000, 1000)
// 	var n, m, x int
// 	fmt.Print("please enter n:")
// 	fmt.Scanln(&n)
// 	fmt.Print("please enter m:")
// 	fmt.Scanln(&m)
// 	fmt.Print("please enter r:")
// 	fmt.Scanln(&x)
// 	for i := 1; i <= n; i++ {
// 		fmt.Print("please enter num:")
// 		fmt.Scanln(&a[i])
// 		b[i] = b[i-1] + a[i]
// 	}
// 	for i := 0; i <= n; i++ {
// 		if b[i]%m == x {
//
// 		}
// 	}
// }

func validate(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func longestPalindromeOne(s string) string {
	size := len(s)
	if size <= 0 {
		return ""
	}
	maxLen := 1
	resStr := s[0:1]
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			subLen := j - i + 1
			if subLen > maxLen && validate(s, i, j) {
				maxLen = subLen
				resStr = s[i : i+subLen]
			}
		}
	}
	return resStr
}

func longestPalindromeTwo(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}
	dp := make([][]bool, 1000)
	for i := range dp {
		dp[i] = make([]bool, 1000)
	}
	for i := range dp {
		dp[i][i] = true
	}
	maxLen := 1
	begin := 0
	for j := 1; j < l; j++ {
		for i := 0; i < j; i++ {
			// 判断s[i:j]是否是回文串
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
				curLen := j - i + 1
				if dp[i][j] && curLen > maxLen {
					// 只要当前是回文串，就记录起始位置和子串的长度
					maxLen = curLen
					begin = i
				}
			}
		}
	}
	return s[begin : begin+maxLen]
}

func longestPalindrome(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}
	maxLen := 1
	resStr := s[0:1]

	for i := 0; i < l-1; i++ {
		oddStr := getSub(s, i, i)
		eveStr := getSub(s, i, i+1)
		curLen, curStr := max1(oddStr, eveStr)
		if curLen > maxLen {
			maxLen = curLen
			resStr = curStr
		}
	}
	return resStr
}

func max1(a, b string) (int, string) {
	if len(a) > len(b) {
		return len(a), a
	}
	return len(b), b
}

func getSub(s string, l, r int) string {
	len := len(s)
	i := l
	j := r
	for i >= 0 && j < len {
		if s[i] == s[j] {
			i--
			j++
		} else {
			break
		}
	}
	return s[i+1 : j-i+1]
}

func main() {
	// fmt.Println(findTheLongestSubstring("eleetminicoworoep"))
	// fmt.Println(Test1([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, 7))
	// PreSumOne()
	// fmt.Println(longestPalindromeOne("aabbaba"))
	// fmt.Println(longestPalindromeTwo("aabbaba"))
	fmt.Println(longestPalindrome("aabbaba"))
}
