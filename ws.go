package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	//"golang.org/x/net/websocket"
)

var broadcast = make(chan Message)
var clients = make(map[*websocket.Conn]bool) //connected clients
var upgrader = websocket.Upgrader{}

type Message struct {
	text      string   `json:"message"`
	email     string   `json:"email"`
	receivers []string `json:"receivers"`
	userName  string   `json:"userName"`
}

func receiveMessages(write http.ResponseWriter, req *http.Request) {
	var err error
	// var jsonString string
	var msg Message
	// var receivers []interface{}
	// var text string
	// var email string
	// var userName string

	ws, err := upgrader.Upgrade(write, req, nil)
	if err != nil {
		fmt.Println(err)
	}

	defer ws.Close()
	for {
		err = ws.ReadJSON(&msg)
		fmt.Println(msg)
		fmt.Println("next line ************")

		// if err = websocket.Message.Receive(ws, &jsonString); err != nil {
		// 	fmt.Println("Can't receive")
		// 	break
		// }

		// fmt.Printf("Initial json = %s \n", jsonString)
		// text, email, userName = getMessageData(jsonString, &receivers)

		// clientsList := make([]string, len(receivers))
		// for i, v := range receivers {
		// 	clientsList[i] = v.(string)
		// 	fmt.Println(clientsList[i])
		// }

		// if err = websocket.Message.Send(ws, reply); err != nil {
		// 	fmt.Println("Can't send")
		// 	break
		// }
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

func main() {
	http.HandleFunc("/", receiveMessages)

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
