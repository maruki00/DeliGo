package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/maruki00/deligo/notifier/models"
	"github.com/maruki00/deligo/notifier/service"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Production usage requires configuration controls
	},
}

type WSMessage struct {
	OrderID     string `json:"order_id"`
	RecipientID string `json:"recipient_id"`
	Message     string `json:"message"`
}

type Hub struct {
	connections map[string]*websocket.Conn
	mu          sync.RWMutex
	svc         service.Service
}

func NewHub(svc service.Service) *Hub {
	return &Hub{
		connections: make(map[string]*websocket.Conn),
		svc:         svc,
	}
}

func (h *Hub) Register(userID string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if oldConn, exists := h.connections[userID]; exists {
		oldConn.Close()
	}
	h.connections[userID] = conn
	log.Printf("User ws session established: %s", userID)
}

func (h *Hub) Unregister(userID string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if conn, exists := h.connections[userID]; exists {
		conn.Close()
		delete(h.connections, userID)
		log.Printf("User ws session terminated: %s", userID)
	}
}

func (h *Hub) RouteMessage(senderID string, payload []byte) {
	var msg WSMessage
	if err := json.Unmarshal(payload, &msg); err != nil {
		log.Printf("Error decoding inbound message structure: %v", err)
		return
	}

	chatModel := models.ChatMessage{
		OrderID:     msg.OrderID,
		SenderID:    senderID,
		RecipientID: msg.RecipientID,
		Message:     msg.Message,
		SentAt:      time.Now(),
	}

	if err := h.svc.SaveChatMessage(&chatModel); err != nil {
		log.Printf("Persistent chat sync mutation execution failure: %v", err)
		return
	}

	h.mu.RLock()
	recipientConn, online := h.connections[msg.RecipientID]
	h.mu.RUnlock()

	if online {
		outboundPayload, _ := json.Marshal(chatModel)
		err := recipientConn.WriteMessage(websocket.TextMessage, outboundPayload)
		if err != nil {
			log.Printf("Failed to pipe instant package transmission down channel to target %s: %v", msg.RecipientID, err)
			h.Unregister(msg.RecipientID)
		}
	} else {
		log.Printf("Recipient %s offline, message cached to persistent storage safely", msg.RecipientID)
	}
}

func (h *Hub) HandleWebSocket(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id query argument required"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Failed to upgrade server protocol parameters: %v", err)
		return
	}

	h.Register(userID, conn)

	defer func() {
		h.Unregister(userID)
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Session closed abruptly or stream error read on %s: %v", userID, err)
			break
		}
		h.RouteMessage(userID, message)
	}
}
