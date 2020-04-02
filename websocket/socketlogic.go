package main

import (
	"fmt"
	"github.com/prometheus/common/log"
	"golang.org/x/net/websocket"
	"net/http"
)

func Echo(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Cant receive")
			break
		}

		fmt.Println("Receive from websocket client:" + reply)

		msg := "Received: " + reply
		fmt.Println("Sending to client: " + msg)

		if err := websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Cant Send,err=:", err)
			break
		}
	}
}

func main() {
	http.Handle("/", websocket.Handler(Echo))

	if err := http.ListenAndServe(":7887", nil); err != nil {
		log.Fatal("ListenAndServer err,err=%v", err)
	}
}
