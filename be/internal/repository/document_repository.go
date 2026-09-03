package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Document struct {
	ID      string
	Title   string
	Content []byte
	OwnerID string
	Version int64
}

type DocumentRepository struct {
	db *pgxpool.Pool
}

func NewDocumentRepository(db *pgxpool.Pool) *DocumentRepository {
	return &DocumentRepository{db: db}
}

func (r *DocumentRepository) CreateDocument(ctx context.Context, doc Document) (Document, error) {
	query := `
		INSERT INTO documents (title, content, owner_id)
		VALUES ($1, $2, $3)
		RETURNING id, title, content, owner_id, version
	`

	var result Document
	err := r.db.QueryRow(ctx, query, doc.Title, doc.Content, doc.OwnerID).
		Scan(&result.ID, &result.Title, &result.Content, &result.OwnerID, &result.Version)
	if err != nil {
		return Document{}, fmt.Errorf("failed to create document: %w", err)
	}

	return result, nil
}

func (r *DocumentRepository) GetDocument(ctx context.Context, id string) (Document, error) {
	query := `
		SELECT id, title, content, owner_id, version
		FROM documents
		WHERE id = $1
	`

	var doc Document
	err := r.db.QueryRow(ctx, query, id).
		Scan(&doc.ID, &doc.Title, &doc.Content, &doc.OwnerID, &doc.Version)
	if err != nil {
		return Document{}, fmt.Errorf("failed to get document: %w", err)
	}

	return doc, nil
}

func (r *DocumentRepository) ListDocuments(ctx context.Context, ownerID string) ([]Document, error) {
	query := `
		SELECT id, title, content, owner_id, version
		FROM documents
		WHERE owner_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query, ownerID)
	if err != nil {
		return nil, fmt.Errorf("failed to list documents: %w", err)
	}
	defer rows.Close()

	var docs []Document
	for rows.Next() {
		var doc Document
		if err := rows.Scan(&doc.ID, &doc.Title, &doc.Content, &doc.OwnerID, &doc.Version); err != nil {
			return nil, err
		}
		docs = append(docs, doc)
	}

	return docs, nil
}

func (r *DocumentRepository) UpdateDocument(ctx context.Context, doc Document) (Document, error) {
	query := `
		UPDATE documents
		SET content = $1, version = version + 1, updated_at = NOW()
		WHERE id = $2
		RETURNING id, title, content, owner_id, version
	`

	var result Document
	err := r.db.QueryRow(ctx, query, doc.Content, doc.ID).
		Scan(&result.ID, &result.Title, &result.Content, &result.OwnerID, &result.Version)
	if err != nil {
		return Document{}, fmt.Errorf("failed to update document: %w", err)
	}

	return result, nil
}

func (r *DocumentRepository) DeleteDocument(ctx context.Context, id string) error {
	query := `DELETE FROM documents WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete document: %w", err)
	}
	return nil
}
