package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "github.com/bbengfort/notes/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type NoteServer struct {
	pb.UnimplementedNoteServiceServer
}

func (n NoteServer) Fetch(ctx context.Context, filter *pb.NoteFilter) (*pb.Notebook, error) {
	return &pb.Notebook{
		Error: nil,
		Notes: []*pb.Note{
			{
				Id:        1,
				Timestamp: "1",
				Author:    "Toto",
				Text:      "Pouet",
				Private:   false,
			},
			{
				Id:        2,
				Timestamp: "2",
				Author:    "Toto",
				Text:      "Pouet2",
				Private:   false,
			},
		},
	}, nil
}

func (n NoteServer) Create(ctx context.Context, note *pb.Note) (*pb.Notebook, error) {
	return &pb.Notebook{
		Error: nil,
		Notes: []*pb.Note{note},
	}, nil
}

func main() {
	endpoint := fmt.Sprintf("localhost:8080")

	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatal(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterNoteServiceServer(grpcServer, &NoteServer{})
	go grpcServer.Serve(lis)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	mopts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = pb.RegisterNoteServiceHandlerFromEndpoint(ctx, mux, endpoint, mopts)
	if err != nil {
		log.Fatal(err)
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	http.ListenAndServe(":9090", mux)
}
