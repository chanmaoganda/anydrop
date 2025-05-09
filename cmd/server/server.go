package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"slices"
	"strconv"

	"github.com/chanmaoganda/anydrop/common"
	pb "github.com/chanmaoganda/anydrop/filetransfer"
	"github.com/hashicorp/mdns"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedFileServiceServer
}

func (s *Server) QueryPlan(ctx context.Context, fileMeta *pb.FileMeta) (*pb.UploadPlan, error) {
	fileHash := fileMeta.GetFileHash()
	plannedChunks := fileMeta.GetPlannedChunks()

	needChunks := make([]int32, 0)

	/// if client requests to remake, all chunks should be resent,
	// and chunks shall be cleaned? maybe not needed, resend chunks will overwrite them
	if fileMeta.GetRemake() {
		for val := range plannedChunks {
			needChunks = append(needChunks, val)
		}

		plan := &pb.UploadPlan{
			FileHash:     fileHash,
			NeededChunks: needChunks,
		}

		return plan, nil
	}

	existingChunks, err := CheckFolderStat(fileHash)
	if err != nil {
		return nil, err
	}

	for value := range plannedChunks {
		if existingChunks[value] {
			continue
		}

		needChunks = append(needChunks, value)
	}

	plan := &pb.UploadPlan{
		FileHash:     fileHash,
		NeededChunks: needChunks,
	}

	return plan, nil
}

func (s *Server) Upload(stream pb.FileService_UploadServer) error {
	fmt.Println("chunk get")

	var fileHash string

	for {
		chunk, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return stream.SendAndClose(&pb.UploadStatus{
				Success: false,
				Message: "Upload Incomplete",
			})
		}

		fileHash = chunk.FileHash

		err = SaveChunk(chunk)

		if err != nil {
			return err
		}
	}

	exists, err := CheckFolderStat(fileHash)
	if err != nil {
		return err
	}

	status := &pb.UploadStatus{
		Success:        true,
		Message:        "Upload Complete",
		ReceivedChunks: int32(len(exists)),
	}

	return stream.SendAndClose(status)
}

func CheckFolderStat(root string) (map[int32]bool, error) {
	_, err := os.Stat(root)

	if err != nil {
		if err := os.Mkdir(root, os.ModePerm); err != nil {
			log.Fatalf("Cannot create dir due to %s\n", err)
		}
	}

	f, err := os.Open(root)

	if err != nil {
		fmt.Println(err)
		return make(map[int32]bool), err
	}

	fileInfo, err := f.ReadDir(-1)

	if err != nil {
		return make(map[int32]bool), err
	}

	chunks := make(map[int32]bool)

	for _, file := range fileInfo {
		index, err := strconv.Atoi(file.Name())
		if err != nil {
			log.Fatalln("err : [", err, "] Cache broke due to ", file.Name())
			return nil, err
		}

		chunks[int32(index)] = true
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

func (s *Server) MakeFile(ctx context.Context, fileMeta *pb.FileMeta) (*pb.UploadStatus, error) {
	existingChunks, err := CheckFolderStat(fileMeta.GetFileHash())
	if err != nil {
		return nil, err
	}

	totalChunks := int32(len(existingChunks))

	if totalChunks != fileMeta.GetPlannedChunks() {
		return &pb.UploadStatus{
			Success:        false,
			Message:        "Missing or Breaking Chunks",
			ReceivedChunks: totalChunks,
		}, nil
	}

	file, err := os.Create(fileMeta.GetFileName())

	if err != nil {
		return nil, err
	}

	sortedChunks := ToSortedList(existingChunks)

	for _, chunkIndex := range sortedChunks {

		chunkPath := fmt.Sprintf("./%s/%d", fileMeta.GetFileHash(), chunkIndex)

		chunkFile, err := os.Open(chunkPath)

		if err != nil {
			return nil, err
		}

		_, err = io.Copy(file, chunkFile)

		if err != nil {
			return nil, err
		}
	}

	log.Printf("File %s Remade", fileMeta.GetFileName())

	return &pb.UploadStatus{
		Success:        true,
		Message:        "Remake successful",
		ReceivedChunks: totalChunks,
	}, nil
}

func ToSortedList(mapper map[int32]bool) []int32 {
	list := make([]int32, 0)
	for key := range mapper {
		list = append(list, key)
	}

	slices.Sort(list)

	return list
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", common.PORT))

	if err != nil {
		log.Fatalln("cannot bind address")
	}

	grpcServer := grpc.NewServer()

	pb.RegisterFileServiceServer(grpcServer, &Server{})

	// mDNS
	host, _ := os.Hostname()

	info := []string{"AnyDrop, cmd tool to transfer files and messages"}

	service, err := mdns.NewMDNSService(host, common.SERVICE_NAME, common.DOMAIN, "", int(common.PORT), nil, info)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("ips ", service.IPs)

	server, err := mdns.NewServer(&mdns.Config{Zone: service})
	if err != nil {
		log.Fatalln(err)
	}

	defer server.Shutdown()
	// end of mDNS

	err = grpcServer.Serve(listen)

	if err != nil {
		log.Fatalln("")
	}
}
