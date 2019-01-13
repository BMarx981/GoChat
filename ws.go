package main

import (
	"encoding/json"
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
	var jsonString string

	for {
		if err = websocket.Message.Receive(ws, &jsonString); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Printf("Initial json = %s \n", jsonString)

		byt := []byte(jsonString)

		var dat map[string]interface{}

		if err := json.Unmarshal(byt, &dat); err != nil {
			panic(err)
		}

		recText := dat["message"].(string)
		fmt.Println(recText)
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
