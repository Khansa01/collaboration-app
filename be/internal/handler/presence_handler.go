package handler

import (
	"context"

	"connectrpc.com/connect"
	presencev1 "github.com/Khansa01/collaboration-app/be/internal/gen/presence/v1"
	"github.com/Khansa01/collaboration-app/be/internal/gen/presence/v1/presencev1connect"
	"github.com/Khansa01/collaboration-app/be/internal/service"
)

type PresenceHandler struct {
	presenceService *service.PresenceService
}

func NewPresenceHandler(presenceService *service.PresenceService) *PresenceHandler {
	return &PresenceHandler{presenceService: presenceService}
}

var _ presencev1connect.PresenceServiceHandler = (*PresenceHandler)(nil)

func (h *PresenceHandler) SyncPresence(
	ctx context.Context,
	stream *connect.BidiStream[presencev1.PresenceRequest, presencev1.PresenceResponse],
) error {
	userID := stream.RequestHeader().Get("X-User-ID")
	if userID == "" {
		return connect.NewError(connect.CodeUnauthenticated, nil)
	}

	return h.presenceService.SyncPresence(ctx, stream, userID)
}
