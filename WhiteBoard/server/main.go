package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	room = make(map[string][]*websocket.Conn)
)

func main() {
	http.HandleFunc("/ws", streamHandler)

	fmt.Println("listening on port 8082")
	log.Fatal(http.ListenAndServe(":8082", nil))

}

func streamHandler(w http.ResponseWriter, r *http.Request) {
	// upgrade this connection to a WebSocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Client Connected")
	room["1"] = append(room["1"], ws)

	// writes new messages indefinitely to our WebSocket connection
	reader(ws)
}

func reader(conn *websocket.Conn) {
	defer func() {
		conn.Close()
	}()

	for {

		n, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		for _, ws := range room["1"] {
			err = ws.WriteMessage(n, msg)

			if err != nil {
				log.Println("write:", err)
				break
			}
		}

	}

}
