package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/bbengfort/notes/v1"
	"google.golang.org/grpc"
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
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:8080"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterNoteServiceServer(grpcServer, &NoteServer{})
	grpcServer.Serve(lis)
}
