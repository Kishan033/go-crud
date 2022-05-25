package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type UserMap struct {

	// defining struct variables
	Name string
	conn *websocket.Conn
}

var userMap = make(map[string]UserMap)

type Event struct {

	// defining struct variables
	Type string
	Data interface{}
}

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		fmt.Println("messageType", messageType)
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		var event Event
		log.Println(string(p))
		parseErr := json.Unmarshal(p, &event)

		if parseErr != nil {

			// if error is not nil
			// print error
			fmt.Println(parseErr)
			return
		}

		switch event.Type {
		case "userJoined":
			defer closeConnection(event.Data.(string), conn)
			if val, ok := userMap[event.Data.(string)]; ok {
				fmt.Println("User exist", val)
			} else {
				userMap[event.Data.(string)] = UserMap{
					event.Data.(string),
					conn,
				}
				broadCastUserList()
			}
		}

		// printing details of
		// decoded data
		fmt.Println("Struct is:", event)

		if err := conn.WriteMessage(messageType, []byte(event.Data.(string))); err != nil {
			log.Println(err)
			return
		}

	}
}

func broadCastUserList() {
	out, err := json.Marshal(Event{
		"onUserChanged",
		userMap,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("string(out)", string(out))
	for key, user := range userMap {
		fmt.Println("Key:", key, "=>", "Element:", user)
		user.conn.WriteMessage(1, []byte(string(out)))
	}
}

func closeConnection(user string, conn *websocket.Conn) {
	fmt.Println("closeConnection called")
	delete(userMap, user)
	broadCastUserList()
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")
	err = ws.WriteMessage(1, []byte("Connection successful"))
	if err != nil {
		log.Println(err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	reader(ws)
}

func WSHandler(r *mux.Router) {
	r.HandleFunc("/ws", wsHandler)
}
