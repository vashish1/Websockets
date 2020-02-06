package main

import (
	"log"
	"net/http"

	socket "github.com/googollee/go-socket.io"
)

func main() {

	server, err := socket.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socket.Socket) {

		log.Println("on connection")

		so.Join("chat")

		so.On("chat message", func(msg string) {
			log.Println("emit:", so.Emit("chat message", msg))
			so.BroadcastTo("chat", "chat message", msg)
		})

		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})

	server.On("error", func(so socket.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
