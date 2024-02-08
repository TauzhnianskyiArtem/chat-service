package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "github.com/TauzhnianskyiArtem/chat-service/pkg/chat_v1"
)

const grpcPort = 9092

type server struct {
	desc.UnimplementedChatV1Server
}

func (s *server) Create(ctx context.Context, req *desc.ChatCreateRequest) (*desc.ChatCreateResponse, error) {
	log.Printf("Usernames: %v", req.GetUsernames())
	log.Printf("Usernames: %v", ctx.Value("test"))

	return &desc.ChatCreateResponse{Id: gofakeit.Int64()}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
