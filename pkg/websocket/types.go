package pkgWebsocket

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Response struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type Message struct {
	IP      string
	Message []byte
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WSClient struct {
	ID   string
	Conn *websocket.Conn
}

type WebSocketServer struct {
	sync.RWMutex
	clients map[string]*websocket.Conn
}
