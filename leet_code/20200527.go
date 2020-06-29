package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		head = l1
		l1 = l1.Next
	} else if l1.Val > l2.Val {
		head = l2
		l2 = l2.Next
	} else {
		head = l1
		l1 = l1.Next
	}
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 == nil && l2 != nil {
		cur.Next = l2
	}
	if l2 == nil && l1 != nil {
		cur.Next = l1
	}

	return head
}

func TestString() {
	// s := "3[a2[c]]"
	// k := 0
	// start := 0
	// res := ""
	// tmpRes := ""
	// for _, v := range s {
	// 	tmpK, err := strconv.Atoi(string(v))
	// 	if err == nil {
	// 		k = tmpK
	// 		continue
	// 	}
	// 	if string(v) != "[" && string(v) != "]" {
	//
	// 		if k != 0 {
	// 			tmpRes += string(v)
	// 		} else {
	// 			res += string(v)
	// 		}
	// 	}
	// 	if string(v) == "[" {
	// 		start++
	// 		continue
	// 	}
	// 	if string(v) == "]" {
	// 		start--
	// 		if start == 0 {
	// 			for k > 0 {
	// 				res += tmpRes
	// 				k--
	// 			}
	// 			tmpRes = ""
	// 		}
	// 	}
	// }
	// fmt.Println(res)
	fmt.Println("]"[0])
}

func main() {
	// head1 := &ListNode{
	// 	Val:  1,
	// 	Next: nil,
	// }
	// head12 := &ListNode{
	// 	Val: 2,
	// }
	// head13 := &ListNode{
	// 	Val: 4,
	// }
	// head2 := &ListNode{
	// 	Val: 1,
	// }
	// head22 := &ListNode{
	// 	Val: 3,
	// }
	// head23 := &ListNode{
	// 	Val: 4,
	// }
	// head1.Next = head12
	// head12.Next = head13
	//
	// head2.Next = head22
	// head22.Next = head23
	// // mergeTwoLists(head1, head2)
	fmt.Println(decodeString("2[leetcode]"))
}

func decodeString(s string) string {
	var stk []string
	ptr := 0
	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '9' {
			digit := getDigit(s, &ptr)
			stk = append(stk, digit)
		} else if cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z' || cur == '[' {
			ptr++
			stk = append(stk, string(cur))
		} else {
			// 遇到了右括弧，开始出栈
			ptr++
			var sub []string
			for stk[len(stk)-1] != "[" {
				sub = append(sub, stk[len(stk)-1])
				stk = stk[:len(stk)-1]
			}
			// 逆转出栈的字符串序列
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-1-i] = sub[len(sub)-1-i], sub[i]
			}
			stk = stk[:len(stk)-1]
			times, _ := strconv.Atoi(stk[len(stk)-1])
			stk = stk[:len(stk)-1]
			stk = append(stk, strings.Repeat(getString(sub), times))
		}
	}
	return getString(stk)
}

func getDigit(s string, ptr *int) string {
	ret := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		ret += string(s[*ptr])
	}
	return ret
}

func getString(s []string) string {
	res := ""
	for _, v := range s {
		res += v
	}
	return res
}
