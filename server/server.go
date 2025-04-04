package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/chanmaoganda/anydrop/filetransfer"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedFileServiceServer
}

func (s *Server) Upload(ctx context.Context, chunk *pb.FileChunk) (*pb.UploadStatus, error) {
	fmt.Println("chunk get")
	return &pb.UploadStatus{
		Success: true,
		Message: "good",
		ReceivedChunks: []int64{chunk.ChunkIndex},
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9000")

	if err != nil {
		log.Fatalln("cannot bind address")
	}

	grpcServer := grpc.NewServer()

	pb.RegisterFileServiceServer(grpcServer, &Server{})

	err = grpcServer.Serve(listen)

	if err != nil {

	}
}