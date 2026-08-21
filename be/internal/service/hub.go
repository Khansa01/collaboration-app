package service

import (
	"sync"

	"connectrpc.com/connect"
	"github.com/Khansa01/collaboration-app/be/internal/gen/collaboration/v1"
)

type Client struct {
	UserID string
	Stream *connect.ServerStream[collaborationv1.OperationResponse]
}

type Hub struct {
	mu    sync.RWMutex
	rooms map[string][]*Client
}

func NewHub() *Hub {
	return &Hub{
		rooms: make(map[string][]*Client),
	}
}

func (h *Hub) Join(docID string, client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.rooms[docID] = append(h.rooms[docID], client)
}

func (h *Hub) Leave(docID, userID string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	clients := h.rooms[docID]
	var updated []*Client
	for _, c := range clients {
		if c.UserID != userID {
			updated = append(updated, c)
		}
	}
	h.rooms[docID] = updated
}

func (h *Hub) Broadcast(docID, senderID string, resp *collaborationv1.OperationResponse) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, c := range h.rooms[docID] {
		if c.UserID != senderID {
			c.Stream.Send(resp)
		}
	}
}
