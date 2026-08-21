package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Khansa01/collaboration-app/be/internal/domain"
	"github.com/Khansa01/collaboration-app/be/internal/gen/auth/v1/authv1connect"
	"github.com/Khansa01/collaboration-app/be/internal/gen/document/v1/documentv1connect"
	"github.com/Khansa01/collaboration-app/be/internal/handler"
	"github.com/Khansa01/collaboration-app/be/internal/repository"
	"github.com/Khansa01/collaboration-app/be/internal/service"
	"github.com/Khansa01/collaboration-app/be/pkg/middleware"
	"github.com/joho/godotenv"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

	// Service
	authService := service.NewAuthService(userRepo)
	docService := service.NewDocumentService(docRepo)

	// Handler
	authHandler := handler.NewAuthHandler(authService)
	docHandler := handler.NewDocumentHandler(docService)

	// Router
	mux := http.NewServeMux()
	path, h := authv1connect.NewAuthServiceHandler(authHandler)
	mux.Handle(path, h)

	docPath, docH := documentv1connect.NewDocumentServiceHandler(docHandler)
	mux.Handle(docPath, docH)

	// Middleware chain
	chain := corsMiddleware(middleware.JWTMiddleware(mux))

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", chain); err != nil {
		log.Fatal(err)
	}
}
