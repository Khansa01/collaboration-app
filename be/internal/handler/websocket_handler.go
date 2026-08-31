package handler

import (
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WSClient struct {
	conn   *websocket.Conn
	send   chan []byte
	docID  string
	userID string
}

type WSHub struct {
	mu    sync.RWMutex
	rooms map[string]map[*WSClient]bool
}

func NewWSHub() *WSHub {
	return &WSHub{
		rooms: make(map[string]map[*WSClient]bool),
	}
}

func (h *WSHub) Join(docID string, client *WSClient) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.rooms[docID] == nil {
		h.rooms[docID] = make(map[*WSClient]bool)
	}
	h.rooms[docID][client] = true
}

func (h *WSHub) Leave(docID string, client *WSClient) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if room, ok := h.rooms[docID]; ok {
		delete(room, client)
		if len(room) == 0 {
			delete(h.rooms, docID)
		}
	}
}

func (h *WSHub) Broadcast(docID string, sender *WSClient, msg []byte) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	for client := range h.rooms[docID] {
		if client != sender {
			select {
			case client.send <- msg:
			default:
				close(client.send)
			}
		}
	}
}

type WebSocketHandler struct {
	hub *WSHub
}

func NewWebSocketHandler(hub *WSHub) *WebSocketHandler {
	return &WebSocketHandler{hub: hub}
}

func (h *WebSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	docID := r.URL.Path[len("/ws/"):]

	// Ambil token dari query param
	tokenStr := r.URL.Query().Get("token")
	userID := "anonymous"

	if tokenStr != "" {
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err == nil && token.Valid {
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				if id, ok := claims["user_id"].(string); ok {
					userID = id
				}
			}
		}
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return
	}

	client := &WSClient{
		conn:   conn,
		send:   make(chan []byte, 256),
		docID:  docID,
		userID: userID,
	}

	h.hub.Join(docID, client)

	// Write pump
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer func() {
			ticker.Stop()
			conn.Close()
		}()
		for {
			select {
			case msg, ok := <-client.send:
				conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
				if !ok {
					conn.WriteMessage(websocket.CloseMessage, []byte{})
					return
				}
				if err := conn.WriteMessage(websocket.BinaryMessage, msg); err != nil {
					return
				}
			case <-ticker.C:
				conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
				if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					return
				}
			}
		}
	}()

	// Read pump
	defer func() {
		h.hub.Leave(docID, client)
		close(client.send)
		conn.Close()
	}()

	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		h.hub.Broadcast(docID, client, msg)
	}
}
