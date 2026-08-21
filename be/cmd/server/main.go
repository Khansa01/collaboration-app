package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Khansa01/collaboration-app/be/internal/domain"
	"github.com/Khansa01/collaboration-app/be/internal/gen/auth/v1/authv1connect"
	"github.com/Khansa01/collaboration-app/be/internal/handler"
	"github.com/Khansa01/collaboration-app/be/internal/repository"
	"github.com/Khansa01/collaboration-app/be/internal/service"
	"github.com/joho/godotenv"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Connect-Protocol-Version, Connect-Timeout-Ms")

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

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	mux := http.NewServeMux()
	path, h := authv1connect.NewAuthServiceHandler(authHandler)
	mux.Handle(path, h)

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", corsMiddleware(mux)); err != nil {
		log.Fatal(err)
	}
}
