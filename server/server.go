package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sort"
	"strconv"

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

    fileName := chunk.FileName
    fileHash := chunk.FileHash

    if err != nil {
        return err
    }

    index := int(chunk.ChunkIndex)

    if !existChunks[index] {
        
        existChunks[index] = true

        err = SaveChunk(chunk)

        if err != nil {
            return err
        }
    }

    for {
        chunk, err = stream.Recv()

        if err != nil {
            err = RemakeFile(existChunks, fileHash, fileName)
            if err != nil {
                fmt.Println(err)
            }
            return stream.SendAndClose(&pb.UploadStatus{
                Success: true,
                Message: "Upload Complete",
            })
        }

        index := int(chunk.ChunkIndex)

        if !existChunks[index] {

            existChunks[index] = true

            err = SaveChunk(chunk)

            if err != nil {
                return err
            }
        }
    }
}

func CheckFolderStat(root string) (map[int]bool, error) {
    _, err := os.Stat(root)

    if err != nil {
        if err := os.Mkdir(root, os.ModePerm); err != nil {
            log.Fatalf("Cannot create dir due to %s\n", err)
        }
    }

    f, err := os.Open(root)

    if err != nil {
        fmt.Println(err)
        return make(map[int]bool), err
    }

    fileInfo, err := f.ReadDir(-1)

    if err != nil {
        return make(map[int]bool), err
    }

    chunks := make(map[int]bool)

    for _, file := range fileInfo {
        index, err := strconv.Atoi(file.Name())
        if err != nil {
            return nil, err
        }

        chunks[index] = true
    }

    return chunks, nil
}

func SaveChunk(chunk *pb.FileChunk) error {
    file, err := os.Create(fmt.Sprintf("./%s/%d", chunk.FileHash, chunk.ChunkIndex))

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

func RemakeFile(chunks map[int]bool, fileHash string, fileName string) error {
    file, err := os.Create(fileName)

    if err != nil {
        return err
    }

    var sortChunks []int
    for index := range chunks {
        sortChunks = append(sortChunks, index)
    }

    sort.Ints(sortChunks)

    log.Printf("Remaking from %d chunks", len(sortChunks))

    for _, chunkIndex := range sortChunks {

        chunkPath := fmt.Sprintf("./%s/%d", fileHash, chunkIndex)

        chunkFile, err := os.Open(chunkPath)

        if err != nil {
            return err
        }

        _, err = io.Copy(file, chunkFile)

        if err != nil {
            return err
        }
    }
    log.Printf("File %s Remade", fileName)
    return file.Close()
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
