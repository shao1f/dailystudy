package main

import (
	"fmt"
	"math"
)

type myStrings string

func (s myStrings) Len() int {
	return len(s)
}

func (s myStrings) char2Int() int {
	l := s.Len()
	// 421
	sum := 0
	for i := 0; i < l; i++ {
		p := l - i - 1
		res := math.Pow(float64(10), float64(p))
		if res > math.MaxInt32 {
			return math.MaxInt32
		}
		sum += int(s[i]-'0') * int(res)
	}
	return sum
}

func myAtoi(str string) int {
	if str == "" {
		return 0
	}
	i := 0
	l := len(str)
	// first := str[0]
	sign := ""
	for ; i < l; i++ {
		if str[i] != ' ' && str[i] != '0' {
			if str[i] == '+' || str[i] == '-' {
				sign = string(str[i])
			} else {
				break
			}
		}
	}
	var num myStrings
	str = str[i:]
	// if str[0] != '+' && str [0] != '-' && (str[0] < '0' && str[0] > '9') {
	//
	// }
	// -42
	if len(str) == 0 {
		return 0
	}

	// if first == '+' || first == '-' {
	// 	sign = string(first)
	// 	str = str[1:]
	// } else if first < '0' || first > '9' {
	// 	return 0
	// }
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			num += myStrings(str[i])
		} else {
			break
		}
	}
	res := num.char2Int()
	switch sign {
	case "-":
		res *= -1
		if res < math.MinInt32 {
			res = math.MinInt32
		}
	case "+":
		fallthrough
	default:
		if res > math.MaxInt32 {
			res = math.MaxInt32
		}
	}
	return res
}

func main() {
	fmt.Println(myAtoi("-000000000000001"))
}
