package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func RecvFromConn(conn net.Conn, exitChannel chan string) {
	// 创建一个接受缓冲
	buf := make([]byte, 1024)

	for {
		// 从套接字中读取数据
		size, err := conn.Read(buf)
		fmt.Println(size)
		if err != nil {
			break
		}
	}
	fmt.Println(string(buf))
	exitChannel <- "read exit"
}

func RecvFromConn1(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	buf := make([]byte, 1024)

	for {
		_, err := conn.Read(buf)
		if err != nil {
			break
		}
	}
	fmt.Println(string(buf))
}

func main() {
	conn, err := net.Dial("tcp", "www.wj-ml.me:443")
	if err != nil {
		fmt.Println(err)
		return
	}

	// exit := make(chan string)
	//
	// go RecvFromConn(conn, exit)
	//
	// time.Sleep(time.Second)
	//
	// conn.Close()
	//
	// fmt.Println(<-exit)

	var wg sync.WaitGroup
	wg.Add(1)
	go RecvFromConn1(conn, &wg)
	time.Sleep(time.Second)
	conn.Close()
	wg.Wait()
	fmt.Println("recv done!")
}
