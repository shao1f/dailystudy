package main

import (
	"fmt"
	"github.com/shao1f/dailystudy/goprogram/rpc_server/model"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// func main() {
// 	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
// 	if err != nil {
// 		log.Fatalf("net dial error,err=%v", err)
// 	}
// 	var reply string
// 	err = client.Call("HelloWorldService.Hello", "world", &reply)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(reply)
// }

// type HelloServiceClient struct {
// 	*rpc.Client
// }
//
// var _ model.HelloServiceInterface = &HelloServiceClient{}
//
// func DialHelloService(netTyp, addr string) (*HelloServiceClient, error) {
// 	conn, err := net.Dial(netTyp, addr)
// 	if err != nil {
// 		fmt.Errorf("dial hello service error,err=%v", err)
// 		return nil, err
// 	}
// 	c := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
// 	return &HelloServiceClient{c}, nil
// }
//
// func (p *HelloServiceClient) Hello(request string, reply *string) error {
// 	return p.Client.Call(model.HelloWorldService+".Hello", request, reply)
// }
//
// func main() {
// 	flag.Parse()
// 	// // fmt.Println(flag.Arg(0))
// 	// // client, err := DialHelloService("tcp", "127.0.0.1:1234")
// 	// if err != nil {
// 	// 	log.Fatal("err dial service,err=%v", err)
// 	// }
// 	// var reply string
// 	// err = client.Hello(flag.Arg(0), &reply)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// fmt.Println(reply)
// 	conn, err := net.Dial("tcp", "127.0.0.1:1234")
// 	if err != nil {
// 		log.Fatal("net dial error,err=%v", err)
// 	}
// 	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
// 	var reply string
// 	err = client.Call(model.HelloWorldService+".Hello", flag.Arg(0), &reply)
// 	if err != nil {
// 		log.Fatalf("client call err,err=%v", err)
// 	}
// 	fmt.Println(reply)
// }

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call(model.HelloWorldService+".Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
