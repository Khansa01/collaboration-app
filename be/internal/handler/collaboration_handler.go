package handler

import (
	"context"

	"connectrpc.com/connect"
	collaborationv1 "github.com/Khansa01/collaboration-app/be/internal/gen/collaboration/v1"
	"github.com/Khansa01/collaboration-app/be/internal/gen/collaboration/v1/collaborationv1connect"
	"github.com/Khansa01/collaboration-app/be/internal/service"
)

type CollaborationHandler struct {
	collabService *service.CollaborationService
}

func NewCollaborationHandler(collabService *service.CollaborationService) *CollaborationHandler {
	return &CollaborationHandler{collabService: collabService}
}

var _ collaborationv1connect.CollaborationServiceHandler = (*CollaborationHandler)(nil)

func (h *CollaborationHandler) SyncDocument(
	ctx context.Context,
	stream *connect.BidiStream[collaborationv1.Operation, collaborationv1.OperationResponse],
) error {
	userID := stream.RequestHeader().Get("X-User-ID")
	if userID == "" {
		return connect.NewError(connect.CodeUnauthenticated, nil)
	}

	return h.collabService.SyncDocument(ctx, stream, userID)
}
