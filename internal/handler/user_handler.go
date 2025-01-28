package handler

import (
	"context"
	"grpc-auth-service/internal/generated/auth"
	"grpc-auth-service/internal/service"
	"log"
)

type UserHandler struct {
	auth.UnimplementedAuthServiceServer
	UserService *service.UserService
}

func NewAuthHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (handler *UserHandler) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	userId, err := handler.UserService.CreateUser(req)
	if err != nil {
		log.Printf("Erro ao criar usuário: %v", err)
		return nil, err
	}

	return &auth.RegisterResponse{
		UserId: userId,
	}, nil
}
