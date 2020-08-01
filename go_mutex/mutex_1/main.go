package main

import (
	"sync"
	"time"
)

var mu1 sync.Mutex
var mu2 sync.Mutex

func mutex1() {
	mu1.Lock()
	mu2.Lock()
	// time.Sleep(time.Second)
	// mu1.Unlock()
}

func mutex2() {
	mu2.Lock()
	mu1.Lock()
	// mu2.Unlock()
}

func main() {
	go mutex1()
	go mutex2()

	time.Sleep(time.Second * 2)
}
