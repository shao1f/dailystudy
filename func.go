package main

import "fmt"

// import (
// 	"flag"
// 	"fmt"
// )
//
// var skillParam = flag.String("skill", "", "skill to perform")
//
// func main() {
// 	flag.Parse()
//
// 	var skill = map[string]func(){
// 		"fire": func() {
// 			fmt.Println("chicken fire")
// 		},
// 		"run": func() {
// 			fmt.Println("soldier run")
// 		},
// 		"fly": func() {
// 			fmt.Println("angel fly")
// 		},
// 	}
//
// 	if f, ok := skill[*skillParam]; ok {
// 		f()
// 	} else {
// 		fmt.Println("skill not found")
// 	}
// }

// // 调用器接口
// type Invoke interface {
// 	// 实现一个call方法
// 	Call(interface{})
// }
//
// type Struct struct {
// }
//
// func (s *Struct) Call(p interface{}) {
// 	fmt.Println("from struct", p)
// }
//
// var invoker Invoke
//
// type FuncCaller func(interface{})
//
// func (f FuncCaller) Call(p interface{}) {
// 	f(p)
// }
//
// func main() {
// 	s := new(Struct)
// 	invoker = s
// 	invoker.Call("hello")
//
// 	invoker = FuncCaller(func(v interface{}) {
// 		fmt.Println("from function", v)
// 	})
// 	invoker.Call("hello")
// }

// func AutoAdd(value int) func() int {
// 	return func() int {
// 		value++
// 		return value
// 	}
// }
//
// func main() {
// 	// str := "hello world"
// 	//
// 	// foo := func() {
// 	// 	str = "hello foo"
// 	// 	fmt.Println(str)
// 	// }
// 	// foo()
// 	// fmt.Println("1", str)
// 	autoAdd1 := AutoAdd(1)
// 	fmt.Println(autoAdd1())
// 	fmt.Println(autoAdd1())
// 	fmt.Printf("%p\n", &autoAdd1)
//
// 	autoAdd2 := AutoAdd(10)
// 	fmt.Println(autoAdd2())
// 	fmt.Println(autoAdd2())
// 	fmt.Printf("%p\n", autoAdd2)
//
// }

// // 玩家生成器
//
// func Player(name string) func() (string, int) {
// 	hp := 150
// 	return func() (string, int) {
// 		return name, hp
// 	}
// }
//
// func main() {
// 	playerAdd := Player("小王")
// 	name, hp := playerAdd()
// 	fmt.Println(name, ":", hp)
// }

// 声明一个解析错误
type ParseError struct {
	Filename string // 文件名
	Line     int    // 行号
}

// 实现error接口，返回错误描述
func (e *ParseError) Error() string {
	return fmt.Sprintf("%s:%d", e.Filename, e.Line)
}

// 创建一些解析错误
func newParseError(filename string, line int) error {
	return &ParseError{filename, line}
}

func main() {
	var e error
	// 创建一个错误实例，包含文件名和行号
	e = newParseError("main", 1)

	// 通过error接口查看错误
	fmt.Println(e.Error())

	// 根据错误接口的具体类型，获取详细的信息
	switch detail := e.(type) {
	case *ParseError:
		fmt.Printf("Filename:%s Line:%d\n", detail.Filename, detail.Line)
	default:
		fmt.Println("other error")
	}
}
