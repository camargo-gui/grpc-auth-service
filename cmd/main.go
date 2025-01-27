package main

import (
	"fmt"
	"grcp-auth-service/config"
	"grcp-auth-service/internal/generated/auth"
	"grcp-auth-service/internal/handler"
	"grcp-auth-service/internal/model"
	"grcp-auth-service/internal/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	port := 50051

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", port, err)
	}

	dbConnection := config.DatabaseConnection();
	err = dbConnection.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	userService := service.NewUserService(dbConnection);
	userHandler := handler.NewAuthHandler(userService);

	server := grpc.NewServer()
	auth.RegisterAuthServiceServer(server, userHandler);

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	log.Printf("Server started on port %d", port)
}
