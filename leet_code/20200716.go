package main

import "fmt"

type Node struct {
	Name  string
	Exist bool
	Next  map[byte]*Node
}

func NewNode() Node {
	return Node{
		Next: make(map[byte]*Node),
	}
}

func (this *Node) Insert(word string) {
	for i := range word {
		val, ok := this.Next[word[i]]
		if !ok {
			this.Next[word[i]] = &Node{
				Name:  string(word[i]),
				Exist: false,
				Next:  make(map[byte]*Node),
			}
			this = this.Next[word[i]]
		} else {
			this = val
		}
	}
	this.Exist = true
}

func (this *Node) Search(word string) bool {
	for i := range word {
		v, ok := this.Next[word[i]]
		if !ok {
			return false
		}
		this = v
	}
	return this.Exist
}

func (this *Node) SearchPrefix(wordPre string) bool {
	for i := range wordPre {
		v, ok := this.Next[wordPre[i]]
		if !ok {
			return false
		}
		this = v
	}
	return true
}

func main() {
	n := NewNode()
	n.Insert("aaabbccca")
	n.Insert("aaa")
	res := n.Search("aaa")
	fmt.Println(res)
}
