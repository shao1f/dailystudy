package main

import (
	"fmt"
	"time"
)

func Process1(a <-chan int) {
	val := <-a
	fmt.Printf("process1 val:%v\n", val)
	return
}

func Process2(b <-chan int) {
	val := <-b
	fmt.Printf("process2 val:%v\n", val)
	return
}

func main() {
	ch := make(chan int)
	go Process1(ch)
	go Process2(ch)
	ch <- 3
	// ch <- 4
	time.Sleep(time.Second)
}
