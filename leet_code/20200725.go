package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Pair1 struct {
	Val int
	Idx int
}

func dailyTemperatures(T []int) []int {
	if len(T) == 0 {
		return nil
	}
	res := make([]int, len(T))
	stak := make([]Pair1, 0, len(T))
	// stak = append(stak,Pair1{
	// 	Val: T[0],
	// 	Idx: 0,
	// })
	// [73, 74, 75, 71, 69, 72, 76, 73]
	for i := 0; i < len(T); i++ {
		for len(stak) != 0 && T[i] > stak[len(stak)-1].Val {
			res[stak[len(stak)-1].Idx] = i - stak[len(stak)-1].Idx
			stak = stak[:len(stak)-1]
		}
		stak = append(stak, Pair1{
			Val: T[i],
			Idx: i,
		})
	}
	for _, v := range stak {
		res[v.Idx] = 0
	}
	return res
}

func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	tmpMap := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}
	stak := make([]string, 0, len(tokens))
	for i := 0; i < len(tokens); i++ {
		if !tmpMap[tokens[i]] {
			stak = append(stak, tokens[i])
		} else {
			f, s := getNums(&stak)
			switch tokens[i] {
			case "+":
				t := f + s
				stak = append(stak, strconv.Itoa(t))
			case "-":
				t := f - s
				stak = append(stak, strconv.Itoa(t))
			case "*":
				t := f * s
				stak = append(stak, strconv.Itoa(t))
			case "/":
				t := f / s
				stak = append(stak, strconv.Itoa(t))
			}
		}
	}
	if len(stak) != 1 {
		return 0
	}
	_, sec := getNums(&stak)
	return sec
}

func getNums(stak *[]string) (first int, second int) {
	// s := int([]byte(stak[len(stak)-1])[0] - '0')
	res := [2]int{}
	for i := 0; i < 2 && len(*stak) > 0; i++ {
		flag := 1
		s := (*stak)[len(*stak)-1]
		*stak = (*stak)[:len(*stak)-1]
		if s[0] == '-' {
			flag = -1
			s = s[1:]
		}
		t, _ := strconv.Atoi(s)
		res[i] = t * flag
	}
	first, second = res[1], res[0]
	return
}

func decodeString(s string) string {
	if len(s) == 0 {
		return ""
	}
	var stak []string
	p := 0
	for p < len(s) {
		cur := s[p]
		if cur >= '0' && cur <= '9' {
			res := GetDigit(s, &p)
			stak = append(stak, res)
		} else if s[p] >= 'a' && s[p] <= 'z' || s[p] >= 'A' && s[p] <= 'Z' || s[p] == '[' {
			stak = append(stak, string(s[p]))
			p++
		} else {
			p++
			// 重新处理，将栈里的字符乘上倍数在入栈
			tmp := []string{}
			for len(stak) != 0 && stak[len(stak)-1] != "[" {
				tmp = append(tmp, stak[len(stak)-1])
				stak = stak[:len(stak)-1]
			}
			// 剔除'['
			stak = stak[:len(stak)-1]
			str := []byte(GetString(tmp))
			// 由于出栈后字符顺序是反的，所以需要将字符串正序
			for i := 0; i < len(str)-2; i++ {
				str[i], str[len(str)-i-1] = str[len(str)-i-1], str[i]
			}
			times, _ := strconv.Atoi(stak[len(stak)-1])
			stak = stak[:len(stak)-1]
			stak = append(stak, strings.Repeat(string(str), times))
		}
	}
	return GetString(stak)
}

func GetString(s []string) string {
	res := ""
	for _, v := range s {
		res += v
	}
	return res
}

func GetDigit(s string, pos *int) string {
	res := ""
	for ; *pos < len(s) && s[*pos] >= '0' && s[*pos] <= '9'; *pos++ {
		res += string(s[*pos])
	}
	return res
}

func Digit(b string) bool {
	i, _ := strconv.Atoi(b)
	return i >= 0 && i <= 9
}

func main() {
	// dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	// r := "13"
	// bs := []byte(r)
	// fmt.Println(bs[0] - '0')
	// strs := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	// evalRPN(strs)
	fmt.Println(decodeString("3[a]2[bc]"))
}
