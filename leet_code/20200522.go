package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

func MakeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

// map里的位置有问题。需要优化
// var indexMap map[int]int
//
// func buildTree1(preorder []int, inorder []int) *TreeNode {
// 	indexMap = make(map[int]int, len(inorder))
// 	for i, v := range inorder {
// 		indexMap[v] = i
// 	}
// 	return Build(preorder, inorder)
// }
//
// func Build(preorder []int, inorder []int) *TreeNode {
// 	if len(preorder) == 0 {
// 		return nil
// 	}
// 	root := MakeNode(preorder[0])
// 	i := indexMap[preorder[0]]
// 	root.Left = Build(preorder[1:len(inorder[:i])+1], inorder[:i])
// 	root.Right = Build(preorder[len(inorder[:i])+1:], inorder[i+1:])
// 	return root
// }

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootNode := MakeNode(preorder[0])
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	rootNode.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	rootNode.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return rootNode
}

func FindSubLen(s string) int {
	tmpMap := make(map[string]int, len(s))
	for i, v := range s {
		tmpMap[string(v)]++
		if tmpMap[string(v)] != 1 {
			return i
		}
	}
	return len(s)
}

func FindSubLenTwo(s string) int {
	tmpMap := make(map[string]int, len(s))
	for i, v := range s {
		tmpMap[string(v)]++
		if tmpMap[string(v)] != 1 {
			return i
		}
	}
	return len(s)
}
func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l <= 0 {
		return 0
	}
	maxLen := 1
	for j := 1; j <= l; j++ {
		for i := 0; i < j; i++ {
			curLen := FindSubLen(s[i:j])
			if curLen > maxLen {
				maxLen = curLen
			}
		}
	}
	return maxLen
}

func lengthOfLongestSubstringTwo(s string) int {
	l := len(s)
	if l <= 0 {
		return 0
	}
	tmpMap := map[byte]int{}
	// -1表示rk指针还没开始移动
	rk, ans := -1, 0
	for i := 0; i < l; i++ {
		if i != 0 {
			// != 0 说明不是第一次移动左指针，需要将map里的当前值删掉
			delete(tmpMap, s[i-1])
		}
		for rk+1 < l && tmpMap[s[rk+1]] == 0 {
			tmpMap[s[rk+1]]++
			rk++
		}
		ans = maax(rk-i+1, ans)
	}
	return ans
}

func maax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// node := MakeNode(3)
	// node1 := MakeNode(9)
	// node2 := MakeNode(20)
	// node3 := MakeNode(15)
	// node4 := MakeNode(7)
	// node.Left = node1
	// node.Right = node2
	// node2.Left = node3
	// node2.Right = node4
	// PreOrder(node)
	// no := buildTree([]int{3, 9, 20, 15, 17}, []int{9, 3, 15, 20, 17})
	// PreOrder(no)
	// // a := []int{1}
	// // fmt.Println(a[1:])
	// buildTree1([]int{3, 9, 20, 15, 17}, []int{9, 3, 15, 20, 17})
	fmt.Println(lengthOfLongestSubstringTwo("abcabcbb"))
	fmt.Println(0 ^ 15)
}
