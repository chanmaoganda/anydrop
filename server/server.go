package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	pb "github.com/chanmaoganda/anydrop/filetransfer"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedFileServiceServer
}

func (s* Server) Upload(stream pb.FileService_UploadServer) error {
	fmt.Println("chunk get")

    chunk, err := stream.Recv()

    if err == io.EOF {
        return stream.SendAndClose(&pb.UploadStatus{
            Success: true,
            Message: "Upload Complete",
        })
    }

    existChunks, err := CheckFolderStat(chunk.FileHash)

    if err != nil {
        return err
    }

    chunkIndex := strconv.Itoa(int(chunk.ChunkIndex))
    if !existChunks[chunkIndex] {
        err = SaveChunk(chunk)
        if err != nil {
            return err
        }
    }

    for {
        chunk, err = stream.Recv()

        if err != nil {
            return stream.SendAndClose(&pb.UploadStatus{
                Success: true,
                Message: "Upload Complete",
            })
        }

        chunkIndex := strconv.Itoa(int(chunk.ChunkIndex))
        if !existChunks[chunkIndex] {
            err = SaveChunk(chunk)
            if err != nil {
                return err
            }
        }
    }
}

func CheckFolderStat(root string) (map[string]bool, error) {
    _, err := os.Stat(root)

    if err != nil {
        if err := os.Mkdir(root, os.ModePerm); err != nil {
            log.Fatalf("Cannot create dir due to %s\n", err)
        }
    }

    f, err := os.Open(root)

    if err != nil {
        fmt.Println(err)
        return make(map[string]bool), err
    }

    fileInfo, err := f.ReadDir(-1)

    if err != nil {
        return make(map[string]bool), err
    }

    chunks := make(map[string]bool)

    for _, file := range fileInfo {
        chunkIndex := strings.TrimPrefix(file.Name(), "chunk")
        chunks[chunkIndex] = true
    }

    return chunks, nil
}

func SaveChunk(chunk *pb.FileChunk) error {
    file, err := os.Create(fmt.Sprintf("./%s/chunk%d", chunk.FileHash, chunk.ChunkIndex))

    if err != nil {
        return err
    }

    _, err = file.Write(chunk.Content)

    if err != nil {
        return err
    }

    file.Close()

    return nil
}

func RemakeFile(chunks map[string]bool, ) {

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
		log.Fatalln("")
	}
}
