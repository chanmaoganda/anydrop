package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/chanmaoganda/anydrop/common"
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

    chunkName := fmt.Sprintf("%s%d", common.CHUNK_PREFIX, chunk.ChunkIndex)

    if !existChunks[chunkName] {
        
        existChunks[chunkName] = true

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

        chunkName := fmt.Sprintf("%s%d", common.CHUNK_PREFIX, chunk.ChunkIndex)

        if !existChunks[chunkName] {

            existChunks[chunkName] = true

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
        chunks[file.Name()] = true
    }

    return chunks, nil
}

func SaveChunk(chunk *pb.FileChunk) error {
    file, err := os.Create(fmt.Sprintf("./%s/%s%d", chunk.FileHash, common.CHUNK_PREFIX, chunk.ChunkIndex))

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

func RemakeFile(chunks map[string]bool, fileHash string, fileName string) error {
    file, err := os.Create(fileName)

    if err != nil {
        return err
    }

    var sortChunks []int
    for chunkName := range chunks {
        indexString := strings.TrimPrefix(chunkName, common.CHUNK_PREFIX)
        index, err := strconv.Atoi(indexString)

        if err != nil {
            return err
        }
        sortChunks = append(sortChunks, index)
    }

    sort.Ints(sortChunks)

    for _, chunkIndex := range sortChunks {

        chunkPath := fmt.Sprintf("./%s/%s%d", fileHash, common.CHUNK_PREFIX, chunkIndex)

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
