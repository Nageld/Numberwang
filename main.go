// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux

package main

import (
	"Numberwang/lobby"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/websocket"
)

var lobbies = make([]*lobby.Lobby, 0)

var port = os.Getenv("PORT")

var upgrader = websocket.Upgrader{} // use default options

// var hub = make(chan []byte, 30)

func echo(w http.ResponseWriter, r *http.Request) {

	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("Upgrade:", err)
		return
	}

	defer c.Close()

	// connections = append(connections, c)
	// id := len(connections)

	l, uid := addToLobby(c)
	go l.Send()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("Read:", err)
			break
		}

		log.Printf("recv: %s", message)

		l.Hub <- []byte("User " + strconv.Itoa(uid) + ": " + string(message))
		// hub <- []byte("User " + strconv.Itoa(id) + ": " + string(message))

	}
}

func addToLobby(conn *websocket.Conn) (*lobby.Lobby, int) {
	lastLobbyID := 0

	for ind, l := range lobbies {
		lastLobbyID = ind
		if l.OccupiedSlots != len(l.Users) {
			l.AddUser(&lobby.User{ind, conn})
			return l, ind
		}
	}

	l := lobby.Lobby{
		"Lobby " + string(lastLobbyID),
		lastLobbyID,
		[5]*lobby.User{},
		make(chan []byte, 30),
		0,
	}

	lobbies = append(lobbies, &l)

	return &l, lastLobbyID
}

func main() {

	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)

	fmt.Println("started")
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {

    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;

    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };

    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        print(JSON.stringify(ws))
        ws.onopen = function(evt) {
            print("OPEN");
            
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };

    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };

    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };

});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
