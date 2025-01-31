package main

import (
	"fmt"
	"grpc-auth-service/config"
	"grpc-auth-service/internal/generated/auth"
	"grpc-auth-service/internal/handler"
	"grpc-auth-service/internal/model"
	"grpc-auth-service/internal/service"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	port := 50051
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", port, err)
	}

	dbConnection := config.DatabaseConnection()
	err = dbConnection.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	secretKey := os.Getenv("JWT_SECRET")
	tokenService := service.NewTokenService(secretKey)
	passwordService := service.NewPasswordService()
	userService := service.NewUserService(dbConnection, passwordService)
	userHandler := handler.NewAuthHandler(userService, passwordService, tokenService)

	server := grpc.NewServer()
	auth.RegisterAuthServiceServer(server, userHandler)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	log.Printf("Server started on port %d", port)
}
