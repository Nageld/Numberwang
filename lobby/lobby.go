package lobby

import "github.com/gorilla/websocket"

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
			go worker(msg, u.Connection)
		}
	}
}

func (l *Lobby) AddUser(newUser *User) {
	for ind, u := range l.Users {
		if u == nil {
			l.Users[ind] = newUser
			l.OccupiedSlots++
		}
	}

}

func worker(message []byte, conn *websocket.Conn) {
	err := conn.WriteMessage(1, message)
	if err != nil {
		// connections = append(connections[:index], connections[index+1:]...)
	}
}
