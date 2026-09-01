package handler

import (
	"encoding/json"
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

type PresenceUser struct {
	UserID   string `json:"userId"`
	Name     string `json:"name"`
	IsOnline bool   `json:"isOnline"`
}

type PresenceMessage struct {
	Type   string         `json:"type"`
	UserID string         `json:"userId"`
	Name   string         `json:"name"`
	Users  []PresenceUser `json:"users"`
}

type PresenceHub struct {
	mu    sync.RWMutex
	rooms map[string]map[string]*PresenceUser
}

func NewPresenceHub() *PresenceHub {
	return &PresenceHub{
		rooms: make(map[string]map[string]*PresenceUser),
	}
}

func (h *PresenceHub) Join(docID string, user *PresenceUser) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.rooms[docID] == nil {
		h.rooms[docID] = make(map[string]*PresenceUser)
	}
	h.rooms[docID][user.UserID] = user
}

func (h *PresenceHub) Leave(docID string, userID string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if room, ok := h.rooms[docID]; ok {
		delete(room, userID)
		if len(room) == 0 {
			delete(h.rooms, docID)
		}
	}
}

func (h *PresenceHub) GetUsers(docID string) []PresenceUser {
	h.mu.RLock()
	defer h.mu.RUnlock()
	var users []PresenceUser
	for _, u := range h.rooms[docID] {
		users = append(users, *u)
	}
	return users
}

func (h *PresenceHub) BroadcastPresence(docID string, clients map[*WSClient]bool, wsHub *WSHub) {
	users := h.GetUsers(docID)
	msg, _ := json.Marshal(PresenceMessage{
		Type:  "presence",
		Users: users,
	})
	wsHub.mu.RLock()
	defer wsHub.mu.RUnlock()
	for client := range clients {
		select {
		case client.send <- msg:
		default:
		}
	}
}

type PresenceWSHandler struct {
	hub   *PresenceHub
	wsHub *WSHub
}

func NewPresenceWSHandler(hub *PresenceHub, wsHub *WSHub) *PresenceWSHandler {
	return &PresenceWSHandler{hub: hub, wsHub: wsHub}
}

func (h *PresenceWSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	docID := r.URL.Path[len("/presence/"):]
	tokenStr := r.URL.Query().Get("token")
	userID := "anonymous"
	name := "Anonymous"

	if tokenStr != "" {
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err == nil && token.Valid {
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				if id, ok := claims["user_id"].(string); ok {
					userID = id
				}
				if n, ok := claims["name"].(string); ok {
					name = n
				}
			}
		}
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("presence upgrade error:", err)
		return
	}
	defer conn.Close()

	user := &PresenceUser{UserID: userID, Name: name, IsOnline: true}
	h.hub.Join(docID, user)
	defer h.hub.Leave(docID, userID)

	// broadcast join
	h.broadcastToRoom(docID, conn)

	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
		h.broadcastToRoom(docID, conn)
	}
}

func (h *PresenceWSHandler) broadcastToRoom(docID string, exclude *websocket.Conn) {
	users := h.hub.GetUsers(docID)
	msg, _ := json.Marshal(PresenceMessage{Type: "presence", Users: users})

	h.wsHub.mu.RLock()
	defer h.wsHub.mu.RUnlock()
	for client := range h.wsHub.rooms[docID] {
		select {
		case client.send <- msg:
		default:
		}
	}
	// juga kirim ke presence conn sendiri
	exclude.WriteMessage(websocket.TextMessage, msg)
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
