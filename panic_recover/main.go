package main

import "errors"

func main() {
	defer func() {
		recover()
	}()
	go func() {
		panic(errors.New("123"))
	}()
}
