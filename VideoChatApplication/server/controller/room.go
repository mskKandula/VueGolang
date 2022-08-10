package controller

import (
	"math/rand"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Participant describes a single entity in a room
type Participant struct {
	Conn *websocket.Conn
	Host bool
}

// Room holds the participants
type Room struct {
	Users map[string][]Participant
	Mutex sync.RWMutex
}

// Creates the room & returns the roomId
func (r *Room) RoomCreation() string {

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, 8)

	rand.Seed(time.Now().UnixNano())

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	roomId := string(b)
	r.Users = make(map[string][]Participant)
	r.Users[roomId] = []Participant{}

	return roomId
}
