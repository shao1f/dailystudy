package main

import "fmt"

func main() {
	str := "abbfbfb"
	for i, v := range str {
		fmt.Println(v)
		fmt.Println(str[:i])
	}
}
