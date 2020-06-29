package main

import (
	"github.com/shao1f/dailystudy/goprogram/rpc_server/model"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// import (
// 	"fmt"
// 	"log"
// 	"net"
// 	"net/rpc"
// )
//
type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello:" + request
	return nil
}

//
// func main() {
// 	rpc.RegisterName("HelloService", new(HelloService))
// 	listener, err := net.Listen("tcp", "127.0.0.1:1234")
// 	if err != nil {
// 		log.Fatal("err,net lister failed,err=%v", err)
// 	}
// 	defer listener.Close()
// 	conn, err := listener.Accept()
// 	if err != nil {
// 		log.Fatal("accept failed,err=%v", err)
// 	}
// 	defer conn.Close()
// 	rpc.ServeConn(conn)
// 	fmt.Println("server succ")
// }

func RegisterServer(svc model.HelloServiceInterface) error {
	return rpc.RegisterName(model.HelloWorldService, svc)
}

func main() {
	// RegisterServer(new(HelloService))
	// listener, err := net.Listen("tcp", "127.0.0.1:1234")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer listener.Close()
	// for {
	// 	conn, err := listener.Accept()
	// 	if err != nil {
	// 		conn.Close()
	// 		fmt.Errorf("get new conn error,err=%v", err)
	// 	}
	// 	defer conn.Close()
	// 	// go rpc.ServeConn(conn)
	// 	go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	// }
	// rpc.RegisterName(model.HelloWorldService,&HelloService{})
	RegisterServer(&HelloService{})
	// listener, err := net.Listen("tcp", "127.0.0.1:1234")
	// if err != nil {
	// 	log.Fatal("err listen:", err)
	// }
	http.HandleFunc("/jsonrpc", func(writer http.ResponseWriter, request *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			writer,
			request.Body,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	http.ListenAndServe(":1234", nil)
}
