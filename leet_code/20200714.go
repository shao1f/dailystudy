package main

import (
	"fmt"
	"reflect"
)

// func addStrings(num1 string, num2 string) string {
// 	flag := 0
// 	c := 0
// 	l1 := len(num1) - 1
// 	l2 := len(num2) - 1
// 	resStr := ""
// 	for l1 >= 0 || l2 >= 0 || flag != 0 {
// 		res := flag
// 		if l1 >= 0 {
// 			res += num1[l1] - '0'
// 			l1--
// 		}
// 		if l2 >= 0 {
// 			res += num2[l2] - '0'
// 			l2--
// 		}
// 		flag = res / 10
// 		resStr += strconv.Itoa(res % 10)
// 	}
// 	return resStr
// }

func main() {
	s := "123"
	for i := 0; i < len(s); i++ {
		fmt.Println((reflect.TypeOf(s[i] - '0').Name()))
		fmt.Println(s[i] - '0')
	}
}
