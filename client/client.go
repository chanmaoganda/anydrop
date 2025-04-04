package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/chanmaoganda/anydrop/filetransfer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("127.0.0.1:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("cannot connect address")
	}

	defer conn.Close()

	client := pb.NewFileServiceClient(conn)

	status, err := client.Upload(context.Background(), &pb.FileChunk{
		Content: []byte{},
		ChunkIndex: 2,
		FileHash: "20121",
	})

	if err != nil {
		log.Fatalln("cannot receive request")
	}

	fmt.Println(status)
}