package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func Echo(ws *websocket.Conn) {
	var err error
	var reply string
	var num int
	num = 0

	for {
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}
		num += 1
		fmt.Printf("%s \n", reply)

		// if err = websocket.Message.Send(ws, reply); err != nil {
		// 	fmt.Println("Can't send")
		// 	break
		// }
	}
}

func main() {
	http.Handle("/", websocket.Handler(Echo))

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
