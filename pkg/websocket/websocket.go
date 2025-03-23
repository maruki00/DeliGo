package pkgWebsocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		clients: make(map[string]*websocket.Conn),
	}
}

func (ws *WebSocketServer) GetClient(clientIp string) *websocket.Conn {
	if ws, ok := ws.clients[clientIp]; ok {
		return ws
	}
	return nil
}

func (ws *WebSocketServer) SetClient(clientIp string, con *websocket.Conn) {
	ws.Lock()
	defer ws.Unlock()
	ws.clients[clientIp] = con
}

func (ws *WebSocketServer) DeleteClient(clientIp string, con *websocket.Conn) {
	ws.Lock()
	defer ws.Unlock()
	delete(ws.clients, clientIp)
}

func (ws *WebSocketServer) Handle(clientIp string, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading:", err)
		return
	}
	ws.clients[clientIp] = conn
}

func (ws *WebSocketServer) handleMessage(clientIp string, msg Response) {
	switch msg.Type {
	case "ping":
		ws.SendResponse(clientIp, "pong", "Pong response")
	case "echo":
		ws.SendResponse(clientIp, "echo", msg.Data)
	default:
		ws.SendResponse(clientIp, "error", "Unknown message type")
	}
}

func (ws *WebSocketServer) SendResponse(clientIp string, msgType string, data string) {
	conn := ws.clients[clientIp]
	resp := Response{
		Type: msgType,
		Data: data,
	}
	err := conn.WriteJSON(resp)
	if err != nil {
		log.Println("Error writing JSON:", err)
	}
}
