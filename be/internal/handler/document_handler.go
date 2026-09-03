package handler

import (
	"context"

	"connectrpc.com/connect"
	documentv1 "github.com/Khansa01/collaboration-app/be/internal/gen/document/v1"
	"github.com/Khansa01/collaboration-app/be/internal/gen/document/v1/documentv1connect"
	"github.com/Khansa01/collaboration-app/be/internal/service"
)

type DocumentHandler struct {
	docService *service.DocumentService
}

func NewDocumentHandler(docService *service.DocumentService) *DocumentHandler {
	return &DocumentHandler{docService: docService}
}

var _ documentv1connect.DocumentServiceHandler = (*DocumentHandler)(nil)

func (h *DocumentHandler) CreateDocument(ctx context.Context, req *connect.Request[documentv1.CreateDocumentRequest]) (*connect.Response[documentv1.CreateDocumentResponse], error) {
	ownerID := req.Header().Get("X-User-ID")
	if ownerID == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, nil)
	}

	doc, err := h.docService.CreateDocument(ctx, req.Msg.Title, ownerID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&documentv1.CreateDocumentResponse{
		Document: &documentv1.Document{
			Id:      doc.ID,
			Title:   doc.Title,
			Content: string(doc.Content),
			OwnerId: doc.OwnerID,
		},
	}), nil
}

func (h *DocumentHandler) GetDocument(ctx context.Context, req *connect.Request[documentv1.GetDocumentRequest]) (*connect.Response[documentv1.GetDocumentResponse], error) {
	doc, err := h.docService.GetDocument(ctx, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&documentv1.GetDocumentResponse{
		Document: &documentv1.Document{
			Id:      doc.ID,
			Title:   doc.Title,
			Content: string(doc.Content),
			OwnerId: doc.OwnerID,
		},
	}), nil
}

func (h *DocumentHandler) ListDocuments(ctx context.Context, req *connect.Request[documentv1.ListDocumentsRequest]) (*connect.Response[documentv1.ListDocumentsResponse], error) {
	ownerID := req.Header().Get("X-User-ID")
	if ownerID == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, nil)
	}

	docs, err := h.docService.ListDocuments(ctx, ownerID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	var result []*documentv1.Document
	for _, doc := range docs {
		result = append(result, &documentv1.Document{
			Id:      doc.ID,
			Title:   doc.Title,
			Content: string(doc.Content),
			OwnerId: doc.OwnerID,
		})
	}

	return connect.NewResponse(&documentv1.ListDocumentsResponse{
		Documents: result,
	}), nil
}

func (h *DocumentHandler) UpdateDocument(ctx context.Context, req *connect.Request[documentv1.UpdateDocumentRequest]) (*connect.Response[documentv1.UpdateDocumentResponse], error) {
	doc, err := h.docService.UpdateDocument(ctx, req.Msg.Id, []byte(req.Msg.Content))
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&documentv1.UpdateDocumentResponse{
		Document: &documentv1.Document{
			Id:      doc.ID,
			Title:   doc.Title,
			Content: string(doc.Content),
			OwnerId: doc.OwnerID,
		},
	}), nil
}

func (h *DocumentHandler) DeleteDocument(ctx context.Context, req *connect.Request[documentv1.GetDocumentRequest]) (*connect.Response[documentv1.GetDocumentResponse], error) {
	ownerID := req.Header().Get("X-User-ID")
	if ownerID == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, nil)
	}

	err := h.docService.DeleteDocument(ctx, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&documentv1.GetDocumentResponse{}), nil
}
