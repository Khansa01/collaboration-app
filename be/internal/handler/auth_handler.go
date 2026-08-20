package handler

import (
	"context"

	"connectrpc.com/connect"
	authv1 "github.com/Khansa01/collaboration-app/be/internal/gen/auth/v1"
	"github.com/Khansa01/collaboration-app/be/internal/gen/auth/v1/authv1connect"
	"github.com/Khansa01/collaboration-app/be/internal/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

var _ authv1connect.AuthServiceHandler = (*AuthHandler)(nil)

func (h *AuthHandler) Register(ctx context.Context, req *connect.Request[authv1.RegisterRequest]) (*connect.Response[authv1.RegisterResponse], error) {
	user, err := h.authService.Register(ctx, req.Msg.Name, req.Msg.Email, req.Msg.Password)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&authv1.RegisterResponse{
		Message: "Welcome " + user.Name + "!",
	}), nil
}

func (h *AuthHandler) Login(ctx context.Context, req *connect.Request[authv1.LoginRequest]) (*connect.Response[authv1.LoginResponse], error) {
	token, err := h.authService.Login(ctx, req.Msg.Email, req.Msg.Password)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	return connect.NewResponse(&authv1.LoginResponse{
		Token: token,
	}), nil
}
