package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var broadcast = make(chan Message)
var clients = make(map[*websocket.Conn]bool) //connected clients
var upgrader = websocket.Upgrader{}

type Message struct {
	text string `json:"message"`
	// email string `json:"email"`
	// receivers []string `json:"receivers"`
	userName string `json:"sender"`
}

func receiveMessages(w http.ResponseWriter, r *http.Request) {
	var err error
	// header := w.Header()
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}

	defer ws.Close()
	clients[ws] = true

	for {
		var msg Message
		i := 1

		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Json read error ")
			fmt.Println(err)
		}
		if i == 1 {
			fmt.Println(msg)
			i += 1
		}
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		broadcast <- msg
	}
}

func handleMessages() {
	i := 0
	for {
		//Grab the next message from the broadcast channel
		byt := <-broadcast
		if i == 1 {
			fmt.Println("in handleMessages")
			fmt.Println(byt)
			i += 1
		}

		//Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(byt)
			if err != nil {
				fmt.Printf("error: %v", err)
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
