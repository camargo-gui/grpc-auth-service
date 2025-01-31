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

	Login(clientConnection, ctx)
}

func Register(client auth.AuthServiceClient, ctx context.Context) {

	req := &auth.RegisterRequest{
		Name:        "John Doe",
		Email:       "gui@gmail.com",
		Password:    "123456",
		Document:    "123456789",
		Phone:       "123456789",
		DateOfBirth: "1990-01-01",
		TenantId:    1,
	}

	resp, err := client.Register(ctx, req)
	log.Printf("User registered with ID: %v, error: %v", resp.GetUserId(), err)
}

func Login(client auth.AuthServiceClient, ctx context.Context) {
	req := &auth.LoginRequest{
		Email:    "teste@gmail.com",
		Password: "123456",
		TenantId: 1,
	}

	resp, err := client.Login(ctx, req)
	log.Printf("User logged in with token: %v, error: %v", resp.GetAccessToken(), err)
}

func Validate(client auth.AuthServiceClient, ctx context.Context) {
	req := &auth.ValidateRequest{
		AccessToken: "",
	}

	resp, err := client.Validate(ctx, req)
	log.Printf("Token is valid for user: %v, error: %v", resp.GetUserId(), err)
}
	
