package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	//"golang.org/x/net/websocket"
)

var broadcast = make(chan string)
var clients = make(map[*websocket.Conn]bool) //connected clients
var upgrader = websocket.Upgrader{}

type Message struct {
	text      string   `json:"message"`
	email     string   `json:"email"`
	receivers []string `json:"receivers"`
	userName  string   `json:"username"`
}

func receiveMessages(write http.ResponseWriter, req *http.Request) {
	var err error
	// var msg string
	// var dat map[string]interface{}

	ws, err := upgrader.Upgrade(write, req, nil)
	if err != nil {
		fmt.Println(err)
	}

	defer ws.Close()
	for {
		text := getJsonMessage(ws)
		fmt.Println(text)
		broadcast <- text
	}
}

func getJsonMessage(sock *websocket.Conn) string {
	_, byt, err := sock.ReadMessage()
	if err != nil {
		fmt.Println(err)
	}

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	text := dat["message"].(string)
	return text
}

func handleMessages() {
	for {
		//Grab the next message from the broadcast channel
		text := <-broadcast
		fmt.Println(text)
		fmt.Println("Reached here")
		//Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(text)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	http.HandleFunc("/", receiveMessages)

	go handleMessages()

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

// func getMessageData(jsonString string, receivers *[]interface{}) (string, string, string) {
// 	byt := []byte(jsonString)
//
// 	var dat map[string]interface{}
//
// 	if err := json.Unmarshal(byt, &dat); err != nil {
// 		panic(err)
// 	}
//
// 	var text = dat["message"].(string)
// 	var email = dat["email"].(string)
// 	var userName = dat["username"].(string)
// 	*receivers = dat["receivers"].([]interface{})
// 	fmt.Println(text + " = text\n" + email + " = email\n" + userName + " = username")
// 	return text, email, userName
// }
