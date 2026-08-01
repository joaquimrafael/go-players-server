package domain

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type playerServerWS struct {
	*websocket.Conn
}

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func NewPlayerServerWS(w http.ResponseWriter, r *http.Request) *playerServerWS {
	conn, err := wsUpgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("problem upgrading connection to WebSockets %v\n", err)
	}

	return &playerServerWS{conn}
}

func (w *playerServerWS) WaitForMsg() string {
	_, msg, err := w.ReadMessage()
	if err != nil {
		log.Printf("error reading from websocket %v\n", err)
	}
	return string(msg)
}

func (w *playerServerWS) Write(message []byte) (int, error) {
	err := w.WriteMessage(websocket.TextMessage, message)
	return len(message), err
}
