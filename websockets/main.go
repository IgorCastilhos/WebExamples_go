package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

// Esse exemplo mostra como trabalhar com Websockets em Go.
// Um servidor simples que ecoa de volta tudo que enviamos.
// Usamos a lib gorilla/websocket

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/echo", func(writer http.ResponseWriter, request *http.Request) {
		conn, _ := upgrader.Upgrade(writer, request, nil) // erro ignorado puramente por simplicidade

		for {
			// Lê mensagem do navegador
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			// Escreve a mensagem no console
			fmt.Printf("%s enviou: %s\n", conn.RemoteAddr(), string(msg))

			// Escreve a mensagem de volta pro navegador
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "websockets/websockets.html")
	})
	http.ListenAndServe(":8080", nil)
}
