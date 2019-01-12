package main

import (
	"fmt"
	"log"
	"net/http"

	// "github.com/gorilla/websocket"
	"golang.org/x/net/websocket"
)

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }
//
// func handler(w http.ResponseWriter, r *http.Request) {
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
//
// 	for {
// 		messageType, reader, err := conn.NextReader()
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		fmt.Println()
// 		writer, err := conn.NextWriter(messageType)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
//
// 		if _, err := io.Copy(writer, reader); err != nil {
// 			return
// 		}
// 		fmt.Println(writer.Write(p))
// 		if err := writer.Close(); err != nil {
// 			return
// 		}
// 	}
// }

func Echo(ws *websocket.Conn) {
	var err error
	var reply string

	// for {
	if err = websocket.Message.Receive(ws, &reply); err != nil {
		fmt.Println("Can't receive")
		// break
	}
	fmt.Println(reply + " is the reply")
	// }

}

func Send(ws *websocket.Conn) {
	var err error
	var reply string
	if err = websocket.Message.Send(ws, reply); err != nil {
		fmt.Println("Can't send")
		// break
	}

	fmt.Println("Send to client: " + reply)
}

func main() {
	http.Handler("/", Echo)
	fmt.Println("Between Funcs")
	http.HandleFunc("/", Send)

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
