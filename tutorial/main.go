package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Connection","upgrade")
	r.Header.Set("Upgrade","websocket")
	r.Header.Set("Sec-Websocket-Version", "13")
	// r.Header.Set("Sec-Websocket-Key","1f76f2njy68yhb75ft7")
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	//.. Use conn to send and receive messages.
	log.Println("Client Connected")
	err = conn.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
	    log.Println(err)
	}

	reader(conn)
}

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}


func runHandlers() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Heya my first web socket project.")
	runHandlers()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
