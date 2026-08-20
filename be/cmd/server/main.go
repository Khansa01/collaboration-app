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

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()

	db, err := domain.NewDB(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Repository
	userRepo := repository.NewUserRepository(db)

	// Service
	authService := service.NewAuthService(userRepo)

	// Handler
	authHandler := handler.NewAuthHandler(authService)

	// Router
	mux := http.NewServeMux()
	path, h := authv1connect.NewAuthServiceHandler(authHandler)
	mux.Handle(path, h)

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
