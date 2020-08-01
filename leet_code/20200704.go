package main

import "fmt"

// func longestPrefix(s string) string {
// 	l := len(s)
// 	if l == 0 || l == 1 {
// 		return ""
// 	}
// 	t := ""
// 	resMap := make(map[string]int)
// 	for i := 0; i < l; i++ {
// 		pre := s[:i]
// 		suf := s[i+1:]
// 		resMap[pre]++
// 		if _, ok := resMap[suf]; ok {
// 			if len(suf) > len(t) {
// 				t = suf
// 			}
// 		}
// 	}
// 	// for j := l - 1; j >= 0; j-- {
// 	// 	tmp := s[j:l]
// 	// 	if _, ok := resMap[s[j:l]]; ok {
// 	// 		if len(tmp) > len(t) {
// 	// 			t = tmp
// 	// 		}
// 	// 	}
// 	// }
// 	return t
// }

func longestPrefix(s string) string {
	l := len(s)
	if l == 0 || l == 1 {
		return ""
	}
	t := ""
	resMap := make(map[string]int)
	for i := 0; i < l; i++ {
		pre := s[:i]
		suf := s[i+1:]
		if _, ok := resMap[pre]; ok {
			if len(pre) > len(t) {
				t = pre
			}
		}
		if _, ok := resMap[suf]; ok {
			if len(suf) > len(t) {
				t = suf
			}
		}
		resMap[pre]++
		resMap[suf]++
	}
	return t
}

func main() {
	fmt.Println(longestPrefix("ababab"))
}
