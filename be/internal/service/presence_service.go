package service

import (
	"context"
	"sync"

	"connectrpc.com/connect"
	presencev1 "github.com/Khansa01/collaboration-app/be/internal/gen/presence/v1"
)

type PresenceHub struct {
	mu    sync.RWMutex
	rooms map[string]map[string]*presencev1.UserPresence
}

func NewPresenceHub() *PresenceHub {
	return &PresenceHub{
		rooms: make(map[string]map[string]*presencev1.UserPresence),
	}
}

type PresenceService struct {
	hub *PresenceHub
}

func NewPresenceService(hub *PresenceHub) *PresenceService {
	return &PresenceService{hub: hub}
}

func (s *PresenceService) SyncPresence(
	ctx context.Context,
	stream *connect.BidiStream[presencev1.PresenceRequest, presencev1.PresenceResponse],
	userID string,
) error {
	var docID string

	defer func() {
		if docID != "" {
			s.hub.mu.Lock()
			if room, ok := s.hub.rooms[docID]; ok {
				delete(room, userID)
			}
			s.hub.mu.Unlock()
		}
	}()

	for {
		req, err := stream.Receive()
		if err != nil {
			return err
		}

		docID = req.DocId

		s.hub.mu.Lock()
		if s.hub.rooms[docID] == nil {
			s.hub.rooms[docID] = make(map[string]*presencev1.UserPresence)
		}
		if req.Presence != nil {
			req.Presence.IsOnline = true
			s.hub.rooms[docID][userID] = req.Presence
		}

		var users []*presencev1.UserPresence
		for _, u := range s.hub.rooms[docID] {
			users = append(users, u)
		}
		s.hub.mu.Unlock()

		if err := stream.Send(&presencev1.PresenceResponse{
			Users: users,
		}); err != nil {
			return err
		}
	}
}
