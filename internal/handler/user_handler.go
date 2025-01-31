package handler

import (
	"context"
	"errors"
	"grpc-auth-service/internal/generated/auth"
	"grpc-auth-service/internal/service"
	"log"
	"time"
)

type UserHandler struct {
	auth.UnimplementedAuthServiceServer
	UserService     *service.UserService
	PasswordService *service.PasswordService
	TokenService    *service.TokenService
}

func NewAuthHandler(
	userService *service.UserService,
	passwordService *service.PasswordService,
	tokenService *service.TokenService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (handler *UserHandler) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	userId, err := handler.UserService.CreateUser(req)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}

	return &auth.RegisterResponse{
		UserId: userId,
	}, nil
}

func (handler *UserHandler) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	user, err := handler.UserService.GetUserByEmail(req.GetEmail(), req.GetTenantId())
	if err != nil {
		return &auth.LoginResponse{
			AccessToken: "",
		}, errors.New("invalid credentials")
	}

	isValidPassword := handler.PasswordService.ComparePassword(user.Password, req.GetPassword())
	if !isValidPassword {
		return &auth.LoginResponse{}, errors.New("invalid credentials")
	}

	token, err := handler.TokenService.GenerateToken(user.ID, time.Hour)

	if err != nil {
		return nil, err
	}

	return &auth.LoginResponse{
		AccessToken: token,
	}, nil
}

func (handler *UserHandler) Validate(ctx context.Context, req *auth.ValidateRequest) (*auth.ValidateResponse, error) {
	userId, err := handler.TokenService.ValidateToken(req.GetAccessToken())

	return &auth.ValidateResponse{
		Valid:  err != nil,
		UserId: userId,
	}, nil
}
