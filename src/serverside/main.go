package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var upgrader = websocket.Upgrader{}

//Message struct
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func main() {

	fs := http.FileServer(http.Dir("../../public"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", handleIncomingConnections)
	go handleAllMessages()
	log.Println("HTTP server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func handleIncomingConnections(w http.ResponseWriter, r *http.Request) {

	log.Println("Going to upgrade")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	clients[ws] = true
	log.Println("ws")
	for {
		var msg Message

		err := ws.ReadJSON(&msg)

		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		broadcast <- msg
		log.Println("Message Recvd")

	}
}

func handleAllMessages() {
	for {
		msg := <-broadcast

		for client := range clients {
			err := client.WriteJSON(msg)

			if err != nil {
				log.Printf("error: %v", err)
				delete(clients, client)
				break
			}

			log.Println("Msg sent")
		}
	}
}
