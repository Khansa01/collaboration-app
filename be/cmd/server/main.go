package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Khansa01/collaboration-app/be/internal/domain"
	"github.com/Khansa01/collaboration-app/be/internal/gen/auth/v1/authv1connect"
	collaborationv1connect "github.com/Khansa01/collaboration-app/be/internal/gen/collaboration/v1/collaborationv1connect"
	"github.com/Khansa01/collaboration-app/be/internal/gen/document/v1/documentv1connect"
	presencev1connect "github.com/Khansa01/collaboration-app/be/internal/gen/presence/v1/presencev1connect"
	"github.com/Khansa01/collaboration-app/be/internal/handler"
	"github.com/Khansa01/collaboration-app/be/internal/repository"
	"github.com/Khansa01/collaboration-app/be/internal/service"
	"github.com/Khansa01/collaboration-app/be/pkg/middleware"
	"github.com/joho/godotenv"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip CORS for WebSocket
		if r.Header.Get("Upgrade") == "websocket" {
			next.ServeHTTP(w, r)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Connect-Protocol-Version, Connect-Timeout-Ms, X-User-ID, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	godotenv.Load()

	ctx := context.Background()

	db, err := domain.NewDB(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Repository
	userRepo := repository.NewUserRepository(db)
	docRepo := repository.NewDocumentRepository(db)

	// Hub
	hub := service.NewHub()
	wsHub := handler.NewWSHub()
	wsHandler := handler.NewWebSocketHandler(wsHub)
	grpcPresenceHub := service.NewPresenceHub()

	// Service
	authService := service.NewAuthService(userRepo)
	docService := service.NewDocumentService(docRepo)
	collabService := service.NewCollaborationService(hub, docRepo)
	presenceService := service.NewPresenceService(grpcPresenceHub)
	// Handler
	authHandler := handler.NewAuthHandler(authService)
	docHandler := handler.NewDocumentHandler(docService)
	collabHandler := handler.NewCollaborationHandler(collabService)
	presenceHandler := handler.NewPresenceHandler(presenceService)

	// Router
	mux := http.NewServeMux()

	path, h := authv1connect.NewAuthServiceHandler(authHandler)
	mux.Handle(path, h)

	docPath, docH := documentv1connect.NewDocumentServiceHandler(docHandler)
	mux.Handle(docPath, docH)

	collabPath, collabH := collaborationv1connect.NewCollaborationServiceHandler(collabHandler)
	mux.Handle(collabPath, collabH)

	presencePath, presenceH := presencev1connect.NewPresenceServiceHandler(presenceHandler)
	mux.Handle(presencePath, presenceH)
	presenceWSHub := handler.NewPresenceHub()
	presenceWSHandler := handler.NewPresenceWSHandler(presenceWSHub, wsHub)
	mux.Handle("/ws/", wsHandler)
	mux.Handle("/presence/", presenceWSHandler)

	// Middleware chain
	chain := corsMiddleware(middleware.JWTMiddleware(mux))

	log.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", chain); err != nil {
		log.Fatal(err)
	}
}
