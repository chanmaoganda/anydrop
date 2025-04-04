package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/chanmaoganda/anydrop/common"
	pb "github.com/chanmaoganda/anydrop/filetransfer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UploadFile(client pb.FileServiceClient, filePath string, fileHash string) error {
	file, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}

	stream, err := client.Upload(context.Background())

	if err != nil {
		return err
	}

	buf := make([]byte, common.CHUNK_SIZE)

	info, err := file.Stat()

	if err != nil {
		return err
	}

	fmt.Printf("Plan to send %d chunks\n", info.Size() / int64(common.CHUNK_SIZE))

	currIndex := 0

	for {
		n, err := file.Read(buf)

		if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatalf("ReadFile Error: %v", err)
        }

		chunk := &pb.FileChunk{
            Content:    buf[:n],
            ChunkIndex: int32(currIndex),
            FileHash:   fileHash,
			FileName: file.Name(),
        }

		if err := stream.Send(chunk); err != nil {
            log.Fatalf("Sending Error: %v", err)
        }

        currIndex += 1
	}

	return nil
}

func main() {
	conn, err := grpc.NewClient("127.0.0.1:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln("cannot connect address")
	}

	defer conn.Close()

	client := pb.NewFileServiceClient(conn)

	err = UploadFile(client, "./test.zip", "aaaaaa")

	if err != nil {
		log.Fatalln(err)
	}
}
