package main

import (
	"context"
	"grpc-auth-service/internal/generated/auth"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	clientConnection := auth.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &auth.RegisterRequest{
		TenantId:    uint32(1),
		Email:       "teste@gmail.com",
		Name:        "Teste",
		Password:    "123456",
		Document:    "12345678901",
		Phone:       "12345678901",
		DateOfBirth: "1990-01-01",
	}

	resp, err := clientConnection.Register(ctx, req)
	if err != nil {
		log.Fatalf("Failed to register user: %v", err)
	}

	log.Println("User registered with ID:", resp.GetUserId())
}
