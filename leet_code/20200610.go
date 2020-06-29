package main

import (
	"fmt"
	"time"
)

// type UpperWriter struct {
// 	io.Writer
// }
//
// func (p *UpperWriter) Write(data []byte) (n int, err error) {
// 	return p.Writer.Write(bytes.ToUpper(data))
// }
// func main() {
// 	test := UpperWriter{
// 		os.Stdout,
// 	}
// 	var str = "abcdEf"
// 	fmt.Fprintln(&test, str)
// }

// type UpString string
//
// func (s UpString) String() string {
// 	return strings.ToUpper(string(s))
// }
//
// func main() {
// 	var test UpString
// 	test = "aEafafsaPsaSASDaaa"
// 	fmt.Println(test)
// }

// var intTmp int
//
// var Total struct {
// 	count int
// 	Mutex sync.Mutex
// }
//
// func Worker(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i < 100; i++ {
// 		// Total.Mutex.Lock()
// 		// Total.count += i
// 		// Total.Mutex.Unlock()
// 		intTmp += i
// 	}
// }
//
// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(5)
// 	go Worker(&wg)
// 	go Worker(&wg)
// 	go Worker(&wg)
// 	go Worker(&wg)
// 	go Worker(&wg)
// 	wg.Wait()
// 	fmt.Println(intTmp)
// }

// var a string
// var done bool
//
// func setup() {
// 	a = "hello, world"
// 	done = true
// }
//
// func main() {
// 	go setup()
// 	for !done {
// 	}
// 	print(a)
// }

// func main() {
// 	a := ""
// 	done := make(chan int, 1)
// 	go func() {
// 		a = "aaaffff"
// 		done <- 1
// 	}()
// 	<-done
// 	fmt.Println(a)
// }

// var done = make(chan bool)
// var msg string
//
// func aGoroutine() {
// 	msg = "hello, world"
// 	<-done
// }
// func main() {
// 	go aGoroutine()
// 	done <- true
// 	println(msg)
// }

// 生产者: 生成 factor 整数倍的序列
func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

// 消费者
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
func main() {
	ch := make(chan int, 64) // 成果队列

	go Producer(3, ch) // 生成 3 的倍数的序列
	go Producer(5, ch) // 生成 5 的倍数的序列
	go Consumer(ch)    // 消费 生成的队列

	// 运行一定时间后退出
	time.Sleep(5 * time.Second)
}
