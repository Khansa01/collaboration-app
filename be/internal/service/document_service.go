package service

import (
	"context"
	"fmt"

	"github.com/Khansa01/collaboration-app/be/internal/repository"
)

type DocumentService struct {
	docRepo *repository.DocumentRepository
}

func NewDocumentService(docRepo *repository.DocumentRepository) *DocumentService {
	return &DocumentService{docRepo: docRepo}
}

func (s *DocumentService) CreateDocument(ctx context.Context, title, ownerID string) (repository.Document, error) {
	if title == "" {
		return repository.Document{}, fmt.Errorf("title cannot be empty")
	}

	doc, err := s.docRepo.CreateDocument(ctx, repository.Document{
		Title:   title,
		Content: []byte("{}"),
		OwnerID: ownerID,
	})
	if err != nil {
		return repository.Document{}, err
	}

	return doc, nil
}

func (s *DocumentService) GetDocument(ctx context.Context, id string) (repository.Document, error) {
	return s.docRepo.GetDocument(ctx, id)
}

func (s *DocumentService) ListDocuments(ctx context.Context, ownerID string) ([]repository.Document, error) {
	return s.docRepo.ListDocuments(ctx, ownerID)
}

func (s *DocumentService) UpdateDocument(ctx context.Context, id string, content []byte) (repository.Document, error) {
	return s.docRepo.UpdateDocument(ctx, repository.Document{
		ID:      id,
		Content: content,
	})
}

func (s *DocumentService) DeleteDocument(ctx context.Context, id string) error {
	return s.docRepo.DeleteDocument(ctx, id)
}
