package documentv1connect

import (
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"

	connect "connectrpc.com/connect"
	v1 "github.com/Khansa01/collaboration-app/be/internal/gen/document/v1"
)

const _ = connect.IsAtLeastVersion1_13_0

const (
	DocumentServiceName                    = "document.v1.DocumentService"
	DocumentServiceCreateDocumentProcedure = "/document.v1.DocumentService/CreateDocument"
	DocumentServiceGetDocumentProcedure    = "/document.v1.DocumentService/GetDocument"
	DocumentServiceListDocumentsProcedure  = "/document.v1.DocumentService/ListDocuments"
	DocumentServiceUpdateDocumentProcedure = "/document.v1.DocumentService/UpdateDocument"
	DocumentServiceDeleteDocumentProcedure = "/document.v1.DocumentService/DeleteDocument"
)

type DocumentServiceClient interface {
	CreateDocument(context.Context, *connect.Request[v1.CreateDocumentRequest]) (*connect.Response[v1.CreateDocumentResponse], error)
	GetDocument(context.Context, *connect.Request[v1.GetDocumentRequest]) (*connect.Response[v1.GetDocumentResponse], error)
	ListDocuments(context.Context, *connect.Request[v1.ListDocumentsRequest]) (*connect.Response[v1.ListDocumentsResponse], error)
	UpdateDocument(context.Context, *connect.Request[v1.UpdateDocumentRequest]) (*connect.Response[v1.UpdateDocumentResponse], error)
	DeleteDocument(context.Context, *connect.Request[v1.GetDocumentRequest]) (*connect.Response[v1.GetDocumentResponse], error)
}

func NewDocumentServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) DocumentServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	documentServiceMethods := v1.File_document_proto.Services().ByName("DocumentService").Methods()
	return &documentServiceClient{
		createDocument: connect.NewClient[v1.CreateDocumentRequest, v1.CreateDocumentResponse](
			httpClient,
			baseURL+DocumentServiceCreateDocumentProcedure,
			connect.WithSchema(documentServiceMethods.ByName("CreateDocument")),
			connect.WithClientOptions(opts...),
		),
		getDocument: connect.NewClient[v1.GetDocumentRequest, v1.GetDocumentResponse](
			httpClient,
			baseURL+DocumentServiceGetDocumentProcedure,
			connect.WithSchema(documentServiceMethods.ByName("GetDocument")),
			connect.WithClientOptions(opts...),
		),
		listDocuments: connect.NewClient[v1.ListDocumentsRequest, v1.ListDocumentsResponse](
			httpClient,
			baseURL+DocumentServiceListDocumentsProcedure,
			connect.WithSchema(documentServiceMethods.ByName("ListDocuments")),
			connect.WithClientOptions(opts...),
		),
		updateDocument: connect.NewClient[v1.UpdateDocumentRequest, v1.UpdateDocumentResponse](
			httpClient,
			baseURL+DocumentServiceUpdateDocumentProcedure,
			connect.WithSchema(documentServiceMethods.ByName("UpdateDocument")),
			connect.WithClientOptions(opts...),
		),
		deleteDocument: connect.NewClient[v1.GetDocumentRequest, v1.GetDocumentResponse](
			httpClient,
			baseURL+DocumentServiceDeleteDocumentProcedure,
			connect.WithSchema(documentServiceMethods.ByName("GetDocument")),
			connect.WithClientOptions(opts...),
		),
	}
}

type documentServiceClient struct {
	createDocument *connect.Client[v1.CreateDocumentRequest, v1.CreateDocumentResponse]
	getDocument    *connect.Client[v1.GetDocumentRequest, v1.GetDocumentResponse]
	listDocuments  *connect.Client[v1.ListDocumentsRequest, v1.ListDocumentsResponse]
	updateDocument *connect.Client[v1.UpdateDocumentRequest, v1.UpdateDocumentResponse]
	deleteDocument *connect.Client[v1.GetDocumentRequest, v1.GetDocumentResponse]
}

func (c *documentServiceClient) CreateDocument(ctx context.Context, req *connect.Request[v1.CreateDocumentRequest]) (*connect.Response[v1.CreateDocumentResponse], error) {
	return c.createDocument.CallUnary(ctx, req)
}
func (c *documentServiceClient) GetDocument(ctx context.Context, req *connect.Request[v1.GetDocumentRequest]) (*connect.Response[v1.GetDocumentResponse], error) {
	return c.getDocument.CallUnary(ctx, req)
}
func (c *documentServiceClient) ListDocuments(ctx context.Context, req *connect.Request[v1.ListDocumentsRequest]) (*connect.Response[v1.ListDocumentsResponse], error) {
	return c.listDocuments.CallUnary(ctx, req)
}
func (c *documentServiceClient) UpdateDocument(ctx context.Context, req *connect.Request[v1.UpdateDocumentRequest]) (*connect.Response[v1.UpdateDocumentResponse], error) {
	return c.updateDocument.CallUnary(ctx, req)
}
func (c *documentServiceClient) DeleteDocument(ctx context.Context, req *connect.Request[v1.GetDocumentRequest]) (*connect.Response[v1.GetDocumentResponse], error) {
	return c.deleteDocument.CallUnary(ctx, req)
}

type DocumentServiceHandler interface {
	CreateDocument(context.Context, *connect.Request[v1.CreateDocumentRequest]) (*connect.Response[v1.CreateDocumentResponse], error)
	GetDocument(context.Context, *connect.Request[v1.GetDocumentRequest]) (*connect.Response[v1.GetDocumentResponse], error)
	ListDocuments(context.Context, *connect.Request[v1.ListDocumentsRequest]) (*connect.Response[v1.ListDocumentsResponse], error)
	UpdateDocument(context.Context, *connect.Request[v1.UpdateDocumentRequest]) (*connect.Response[v1.UpdateDocumentResponse], error)
	DeleteDocument(context.Context, *connect.Request[v1.GetDocumentRequest]) (*connect.Response[v1.GetDocumentResponse], error)
}

func NewDocumentServiceHandler(svc DocumentServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	documentServiceMethods := v1.File_document_proto.Services().ByName("DocumentService").Methods()
	documentServiceCreateDocumentHandler := connect.NewUnaryHandler(
		DocumentServiceCreateDocumentProcedure, svc.CreateDocument,
		connect.WithSchema(documentServiceMethods.ByName("CreateDocument")),
		connect.WithHandlerOptions(opts...),
	)
	documentServiceGetDocumentHandler := connect.NewUnaryHandler(
		DocumentServiceGetDocumentProcedure, svc.GetDocument,
		connect.WithSchema(documentServiceMethods.ByName("GetDocument")),
		connect.WithHandlerOptions(opts...),
	)
	documentServiceListDocumentsHandler := connect.NewUnaryHandler(
		DocumentServiceListDocumentsProcedure, svc.ListDocuments,
		connect.WithSchema(documentServiceMethods.ByName("ListDocuments")),
		connect.WithHandlerOptions(opts...),
	)
	documentServiceUpdateDocumentHandler := connect.NewUnaryHandler(
		DocumentServiceUpdateDocumentProcedure, svc.UpdateDocument,
		connect.WithSchema(documentServiceMethods.ByName("UpdateDocument")),
		connect.WithHandlerOptions(opts...),
	)
	documentServiceDeleteDocumentHandler := connect.NewUnaryHandler(
		DocumentServiceDeleteDocumentProcedure, svc.DeleteDocument,
		connect.WithSchema(documentServiceMethods.ByName("DeleteDocument")),
		connect.WithHandlerOptions(opts...),
	)
	return "/document.v1.DocumentService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case DocumentServiceCreateDocumentProcedure:
			documentServiceCreateDocumentHandler.ServeHTTP(w, r)
		case DocumentServiceGetDocumentProcedure:
			documentServiceGetDocumentHandler.ServeHTTP(w, r)
		case DocumentServiceListDocumentsProcedure:
			documentServiceListDocumentsHandler.ServeHTTP(w, r)
		case DocumentServiceUpdateDocumentProcedure:
			documentServiceUpdateDocumentHandler.ServeHTTP(w, r)
		case DocumentServiceDeleteDocumentProcedure:
			documentServiceDeleteDocumentHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

type UnimplementedDocumentServiceHandler struct{}

func (UnimplementedDocumentServiceHandler) CreateDocument(context.Context, *connect.Request[v1.CreateDocumentRequest]) (*connect.Response[v1.CreateDocumentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("document.v1.DocumentService.CreateDocument is not implemented"))
}
func (UnimplementedDocumentServiceHandler) GetDocument(context.Context, *connect.Request[v1.GetDocumentRequest]) (*connect.Response[v1.GetDocumentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("document.v1.DocumentService.GetDocument is not implemented"))
}
func (UnimplementedDocumentServiceHandler) ListDocuments(context.Context, *connect.Request[v1.ListDocumentsRequest]) (*connect.Response[v1.ListDocumentsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("document.v1.DocumentService.ListDocuments is not implemented"))
}
func (UnimplementedDocumentServiceHandler) UpdateDocument(context.Context, *connect.Request[v1.UpdateDocumentRequest]) (*connect.Response[v1.UpdateDocumentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("document.v1.DocumentService.UpdateDocument is not implemented"))
}
func (UnimplementedDocumentServiceHandler) DeleteDocument(context.Context, *connect.Request[v1.GetDocumentRequest]) (*connect.Response[v1.GetDocumentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("document.v1.DocumentService.DeleteDocument is not implemented"))
}
