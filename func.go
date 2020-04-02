package main

import (
	"fmt"
	"sort"
)

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

// // 声明一个解析错误
// type ParseError struct {
// 	Filename string // 文件名
// 	Line     int    // 行号
// }
//
// // 实现error接口，返回错误描述
// func (e *ParseError) Error() string {
// 	return fmt.Sprintf("%s:%d", e.Filename, e.Line)
// }
//
// // 创建一些解析错误
// func newParseError(filename string, line int) error {
// 	return &ParseError{filename, line}
// }
//
// func main() {
// 	var e error
// 	// 创建一个错误实例，包含文件名和行号
// 	e = newParseError("main", 1)
//
// 	// 通过error接口查看错误
// 	fmt.Println(e.Error())
//
// 	// 根据错误接口的具体类型，获取详细的信息
// 	switch detail := e.(type) {
// 	case *ParseError:
// 		fmt.Printf("Filename:%s Line:%d\n", detail.Filename, detail.Line)
// 	default:
// 		fmt.Println("other error")
// 	}
// }

// // 定义一个数据写入器
// type DataWriter interface {
// 	WriteData(data interface{}) error
// 	CanWrite() bool
// }
//
// // 定义文件结构，用于实现DataWriter
// type file struct {
// }
//
// // 实现DataWriter接口的WriteData方法
// func (f *file) WriteData(data interface{}) error {
// 	// 模拟写入数据
// 	fmt.Println("WriteData:", data)
// 	return nil
// }
//
// // 实现CanWrite接口
// func (f *file) CanWrite() bool {
// 	return true
// }
//
// func main() {
// 	// 实例化file
// 	f := new(file)
//
// 	// 声明一个DataWriter的接口
// 	var writer DataWriter
//
// 	// 将接口赋值f，*file类型
// 	writer = f
//
// 	// 使用DataWriter的接口进行数据写入
// 	err := writer.WriteData("data")
// 	if err != nil {
// 		log.Fatalf("Write data error,err=%f", err)
// 	}
// }

// // 声明日志写入器接口
// type LogWriter interface {
// 	Write(data interface{}) error
// }
//
// // 日志器
// type Logger struct {
// 	// 这个日志用到的日志写入器
// 	writerList []LogWriter
// }
//
// // 注册一个日志写入器
// func (l *Logger) RegisterWriter(writer LogWriter) {
// 	l.writerList = append(l.writerList, writer)
// }
//
// // 将一个data类型的数据写入日志
// func (l *Logger) Log(data interface{}) {
// 	// 遍历所有的注册的写入器
// 	for _, writer := range l.writerList {
// 		// 将日志输出到每一个写入器中
// 		_ = writer.Write(data)
// 	}
// }
//
// // 创建日志器实例
// func NewLogger() *Logger {
// 	return &Logger{}
// }
//
// // 声明文件写入器
// type fileWriter struct {
// 	file *os.File
// }
//
// // 设置文件写入器写入的文件名
// func (f *fileWriter) SetFile(fileName string) (err error) {
// 	// 如果文件已经打开，关闭前一个文件
// 	if f.file != nil {
// 		_ = f.file.Close()
// 	}
// 	// 创建一个文件并保存句柄
// 	f.file, err = os.Create(fileName)
//
// 	// 如果创建的过程中出错，返回错误
// 	return err
// }
//
// // 实现Logger的Writer()方法
// func (f *fileWriter) Write(data interface{}) error {
// 	// 日志文件可能没有创建成功
// 	if f.file == nil {
// 		// 日志文件没准备好
// 		return errors.New("file not created")
// 	}
//
// 	// 将数据序列化为字符串
// 	str := fmt.Sprintf("%v\n", data)
//
// 	// 将数据以字节数组写入文件中
// 	_, err := f.file.Write([]byte(str))
//
// 	return err
// }
//
// // 创建文件写入器实例
// func CreateFileWriter() *fileWriter {
// 	return &fileWriter{}
// }
//
// // 命令行写入器
// type consoleWriter struct {
// }
//
// func (f *consoleWriter) Write(data interface{}) error {
// 	// 将数据序列化成字符串
// 	str := fmt.Sprintf("%v\n", data)
// 	// 将数据以字节数组的格式写入到命令行
// 	_, err := os.Stdout.Write([]byte(str))
//
// 	return err
// }
//
// func CreateConsoleWriter() *consoleWriter {
// 	return &consoleWriter{}
// }
//
// func createLogger() *Logger {
// 	//创建日志器
// 	l := NewLogger()
//
// 	// 创建命令行写入器
// 	cw := CreateConsoleWriter()
//
// 	// 注册命令行写入器
// 	l.RegisterWriter(cw)
//
// 	// 创建文件写入器
// 	f := CreateFileWriter()
//
// 	// 设置文件
// 	if err := f.SetFile("log.log"); err != nil {
// 		fmt.Println(err)
// 	}
//
// 	// 注册文件写入器
// 	l.RegisterWriter(f)
//
// 	return l
// }
//
// func main() {
// 	// 准备日志器
// 	l := createLogger()
//
// 	// 写一个日志
// 	l.Log("main.go")
// }

// 将[]string定义为MySTringList类型
type MyStringList []string

// 实现sort.Interface接口获取元素数量方法
func (m MyStringList) Len() int {
	return len(m)
}

// 实现sort.Interface的比较元素方法
func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

// 实现sort.Interface的交换元素方法
func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {
	// 准备一个内容顺序被打乱的字符串切片
	names := MyStringList{
		"5.Penta Kill",
		"3.Triple Kill",
		"1.First Blood",
		"4.Quadra Kill",
		"2.Double Kill",
	}

	// 使用sort包进行排序
	sort.Sort(names)

	// 遍历打印结果
	for _, v := range names {
		fmt.Println(v)
	}
}
