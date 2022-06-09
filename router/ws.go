package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type UserMap struct {
	Name string
	conn *websocket.Conn
}

var userMap = make(map[string]UserMap)

type Event struct {
	Type string
	Data interface{}
}

type EventSendMessage struct {
	Chat    string
	From    string
	To      string
	Message string
}

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		var event Event
		log.Println(string(p))
		parseErr := json.Unmarshal(p, &event)

		if parseErr != nil {
			return
		}

		isHandshake := true

		switch event.Type {
		case "userJoined":
			{
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
				isHandshake = false
			}
		case "onMessageSend":
			{
				var data EventSendMessage
				mapstructure.Decode(event.Data, &data)
				if receiver, ok := userMap[data.To]; ok {
					fmt.Println("User exist", receiver)
					if sender, ok := userMap[data.From]; ok {
						fmt.Println("User exist", sender)

						message, err := json.Marshal(Event{
							"onMessageReceived",
							map[string]string{
								"to":      data.To,
								"from":    data.From,
								"message": data.Message,
								"chat":    data.To,
							},
						})
						if err != nil {
							panic(err)
						}
						receiver.conn.WriteMessage(1, []byte(message))
						sender.conn.WriteMessage(1, []byte(message))
					}
				}
				isHandshake = false
			}
		}

		fmt.Println("Struct is:", event)
		if isHandshake {
			if err := conn.WriteMessage(messageType, []byte(event.Data.(string))); err != nil {
				log.Println(err)
				return
			}
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

func WSHandler() *mux.Router {
	mux := mux.NewRouter().StrictSlash(false)
	mux.HandleFunc("/ws", wsHandler)
	return mux
}
