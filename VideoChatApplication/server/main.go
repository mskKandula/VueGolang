package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mskKandula/VideoChatApp/controller"
)

func main() {
	// var room controller.Room
	// room.Iniit()

	http.HandleFunc("/create", controller.CreateRoom)
	http.HandleFunc("/join", controller.JoinRoom)
	fmt.Println("Hello world")

	log.Fatal(http.ListenAndServe(":8000", nil))
}
