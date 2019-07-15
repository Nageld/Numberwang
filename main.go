// +build linux

package main

import (
	"Numberwang/lobutils"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var lobbies = make([]*lobutils.Lobby, 0)

var port = os.Getenv("PORT")

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
} // use default options

func echo(w http.ResponseWriter, r *http.Request) {

	// testTemplate.Execute(w, lobbies)

	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("Upgrade:", err)
		return
	}

	defer c.Close()

	l, uid := addToLobby(c)
	fmt.Println("Connection opened in \nLobby: " + strconv.Itoa(l.ID) + "\nUser: " + strconv.Itoa(uid))
	fmt.Println("-----------------------------")

	for _, l := range lobbies {
		fmt.Println(l)
	}

	go l.Send()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("Read:", err)
			break
		}

		log.Printf("Lobby: %d User: %d \nSent: %s", l.ID, uid, message)

		l.Hub <- []byte("User " + strconv.Itoa(uid) + ": " + string(message))

	}
}

func addToLobby(conn *websocket.Conn) (*lobutils.Lobby, int) {
	lastLobbyID := 0

	// Go through lobbies and see if there is an empty user slot in any existing lobbies
	for ind, l := range lobbies {
		lastLobbyID = ind
		if l.OccupiedSlots != len(l.Users) {
			uid := l.AddUser(conn)
			return l, uid
		}
	}

	// If every lobby is full or there are no lobbies
	l := lobutils.Lobby{
		"Lobby " + string(lastLobbyID),
		len(lobbies),
		uuid.New(),
		[5]*lobutils.User{},
		make(chan []byte, 30),
		0,
	}

	uid := l.AddUser(conn)
	go l.Ping()

	lobbies = append(lobbies, &l)

	return &l, uid
}

func main() {

	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	// http.HandleFunc("/potato", testHome)
	http.HandleFunc("/", home)

	fmt.Println("main.go Started")

	log.Fatal(http.ListenAndServe(":"+port, nil))

}

func home(w http.ResponseWriter, r *http.Request) {
	// homeTemplate.Execute(w, "wss://"+r.Host+"/echo")
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo") // Use this line when debugging locally

}

func testHome(w http.ResponseWriter, r *http.Request) {
	testTemplate.Execute(w, lobbies)
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
			if (evt.data != "") {
				print("RESPONSE: " + evt.data);
			}

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

var testTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
</head>
<body>
{{.}}
</body>
</html>
`))
