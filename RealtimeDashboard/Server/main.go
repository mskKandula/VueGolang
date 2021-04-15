package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)


type IntegerArray struct{
	Intarr [] int 
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}


func reader(conn *websocket.Conn) {

	for {
		rand.Seed(time.Now().UnixNano())
     
		i :=&IntegerArray{
			
			Intarr : rand.Perm(12),
		}

		if err := conn.WriteJSON(i); err != nil {
			log.Println(err)
			return
		}
		
		time.Sleep(3 * time.Second)
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")

	err = ws.WriteMessage(1, []byte("Hi Client!"))

	if err != nil {
		log.Println(err)
	}
	
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/barws", wsEndpoint)
}

func main() {
	fmt.Println("Hello World")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}