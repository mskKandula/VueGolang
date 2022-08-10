package controller

import (
	"encoding/json"
	"net/http"
)

var (
	room Room
)

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	roomID := room.RoomCreation()

	json.NewEncoder(w).Encode(map[string]string{
		"roomId": roomID,
	})
}

func JoinRoom(w http.ResponseWriter, r *http.Request) {}
