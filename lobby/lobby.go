package lobby

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

type Lobby struct {
	Name          string
	ID            int
	Users         [5]*User
	Hub           chan []byte
	OccupiedSlots int
}

type User struct {
	ID         int
	Connection *websocket.Conn
}

func (l *Lobby) Send() {
	for msg := range l.Hub {
		for _, u := range l.Users {
			if u != nil {
				_ = u
				go worker(msg, u, l)
			}

		}
	}
}

func (l *Lobby) Ping() {
	for {
		time.Sleep(5 * time.Second)

		for _, u := range l.Users {
			if u != nil {
				_ = u
				go worker([]byte(""), u, l)
			}

		}
	}
}

func (l *Lobby) AddUser(conn *websocket.Conn) int {
	newUser := User{-1, conn}

	for ind, u := range l.Users {
		if u == nil {
			newUser.ID = ind
			l.Users[ind] = &newUser
			l.OccupiedSlots++
			return ind
		}
	}

	return 0

}

func worker(message []byte, user *User, lobby *Lobby) {
	err := user.Connection.WriteMessage(1, message)
	if err != nil {
		lobby.Users[user.ID] = nil
		fmt.Println("Removed: " + "\nUser: " + strconv.Itoa(user.ID) + "\nLobby: " + strconv.Itoa(lobby.ID))
		lobby.OccupiedSlots--
	}

}
