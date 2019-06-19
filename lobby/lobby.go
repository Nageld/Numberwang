package lobby

import "github.com/gorilla/websocket"

type Lobby struct {
	name          string
	id            int
	users         [5]*user
	hub           chan []byte
	occupiedSlots int
}

type user struct {
	uid        int
	connection *websocket.Conn
}

func (l *Lobby) send() {
	for msg := range l.hub {
		for _, u := range l.users {
			go worker(msg, u.connection)
		}
	}
}

func (l *Lobby) addUser(newUser *user) {
	for ind, user := range l.users {
		if user == nil {
			l.users[ind] = newUser
			l.occupiedSlots++
		}
	}

}

func worker(message []byte, conn *websocket.Conn) {
	err := conn.WriteMessage(1, message)
	if err != nil {
		// connections = append(connections[:index], connections[index+1:]...)
	}
}
