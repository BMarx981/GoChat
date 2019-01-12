package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

var broadcast = make(chan Message)

type Message struct {
	text      string   `json:"message"`
	email     string   `json:"email"`
	receivers []string `json:"receivers"`
	userName  string   `json:"userName"`
}

func Receive(ws *websocket.Conn) {
	var err error
	var json string

	for {
		if err = websocket.Message.Receive(ws, &json); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Printf("%s \n", json)

		// if err = websocket.Message.Send(ws, reply); err != nil {
		// 	fmt.Println("Can't send")
		// 	break
		// }
	}
}

func main() {
	http.Handle("/", websocket.Handler(Receive))

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
