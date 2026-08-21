package service

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	collaborationv1 "github.com/Khansa01/collaboration-app/be/internal/gen/collaboration/v1"
	"github.com/Khansa01/collaboration-app/be/internal/repository"
)

type CollaborationService struct {
	hub     *Hub
	docRepo *repository.DocumentRepository
}

func NewCollaborationService(hub *Hub, docRepo *repository.DocumentRepository) *CollaborationService {
	return &CollaborationService{hub: hub, docRepo: docRepo}
}

func (s *CollaborationService) SyncDocument(
	ctx context.Context,
	stream *connect.BidiStream[collaborationv1.Operation, collaborationv1.OperationResponse],
	userID string,
) error {
	var docID string

	for {
		op, err := stream.Receive()
		if err != nil {
			if docID != "" {
				s.hub.Leave(docID, userID)
			}
			return err
		}

		if docID == "" {
			docID = op.DocId
			client := &Client{
				UserID: userID,
				Stream: nil,
			}
			_ = client
		}

		_, err = s.docRepo.UpdateDocument(ctx, repository.Document{
			ID:      op.DocId,
			Content: op.Payload,
		})
		if err != nil {
			return fmt.Errorf("failed to update document: %w", err)
		}

		resp := &collaborationv1.OperationResponse{
			Success: true,
			Version: op.Version + 1,
		}

		if err := stream.Send(resp); err != nil {
			return err
		}

		s.hub.Broadcast(op.DocId, userID, resp)
	}
}
