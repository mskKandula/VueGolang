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
	CheckOrigin:     func(r *http.Request) bool { return true },
}

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

	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
		return
	}

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

		// cmd := exec.Command("ffmpeg", "-v", "verbose", "-i", msg,
		// 	"-vf", "scale=1920:1080", "-vcodec", "libx264", "-r", "25", "-b:v",
		// 	"1000000", "-crf", "31", "-acodec", "aac", "-sc_threshold", "0", "-f", "hls",
		// 	"-hls_time", "5", "-segment_time", "5", "-hls_list_size", "5", "stream.m3u8")

		// err = cmd.Run()

		// if err != nil {
		// 	fmt.Println("75", err)
		// }

		err = conn.WriteMessage(n, msg)

		if err != nil {
			log.Println("write:", err)
			break
		}
	}

}
