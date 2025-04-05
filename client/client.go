package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"time"

	"github.com/chanmaoganda/anydrop/common"
	pb "github.com/chanmaoganda/anydrop/filetransfer"
	"github.com/hashicorp/mdns"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func QueryFilePlan(client pb.FileServiceClient, filePath string, plannedChunks int32, needRemake bool) (*pb.UploadPlan, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}

	checkSum, err := common.CheckSumSha256(filePath)

	if err != nil {
		return nil, err
	}

	return client.QueryPlan(context.Background(), &pb.FileMeta{
		FileHash: checkSum,
		FileName: file.Name(),
		PlannedChunks: plannedChunks,
		Remake: needRemake,
	})
}

func UploadFile(client pb.FileServiceClient, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}

	info, err := file.Stat()
	if err != nil {
		return err
	}
	plannedChunks := int32(math.Ceil(float64(info.Size()) / float64(common.CHUNK_SIZE)))

	var checkSum string
	var needRemake bool = false

	for attempt := 0; attempt < int(common.RETRY_TIMES); attempt += 1 {
		var innerRemakeFlag bool

		// first attempt should be controlled by user desire, while in other attempts, remake is false
		if attempt == 0 {
			innerRemakeFlag = needRemake
		} else {
			innerRemakeFlag = false
		}

		plan, err := QueryFilePlan(client, filePath, plannedChunks, innerRemakeFlag)
		if err != nil {
			return err
		}

		checkSum = plan.GetFileHash()

		log.Println("FileHash", plan.FileHash, "needed, ", plan.NeededChunks)

		status, err := UploadChunk(client, plan, file, plan.GetNeededChunks())

		if err != nil {
			return err
		}

		if status.GetSuccess() {
			break
		}
	}

	log.Printf("Chunks Successfully Sent\n")

	status, err := client.MakeFile(context.Background(), &pb.FileMeta{
		FileName: file.Name(),
		FileHash: checkSum,
		PlannedChunks: plannedChunks,
	})

	if err != nil {
		return err
	}

	if !status.GetSuccess() {
		log.Println(status.GetMessage())
	}

	return nil
}

func UploadChunk(client pb.FileServiceClient, plan *pb.UploadPlan, file *os.File, needed []int32) (*pb.UploadStatus, error) {
	stream, err := client.Upload(context.Background())

	if err != nil {
		return nil, err
	}

	buf := make([]byte, common.CHUNK_SIZE)

	if len(plan.NeededChunks) == 0 {
		plan.NeededChunks = append(plan.NeededChunks, 0)
	}

	for _, chunk := range plan.NeededChunks {
		offset := int64(chunk) * common.CHUNK_SIZE

		if _, err := file.Seek(offset, io.SeekStart); err != nil {
			return nil, err
		}

		n, err := file.Read(buf)

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("ReadFile Error: %v\n", err)
		}

		chunk := &pb.FileChunk{
			Content:    buf[:n],
			ChunkIndex: chunk,
			FileHash:   plan.GetFileHash(),
			FileName:   file.Name(),
		}

		if err := stream.Send(chunk); err != nil {
			log.Fatalf("Sending Error: %v", err)
		}

		if err != nil {
			log.Fatalf("Receiving Error %s\n", err)
		}
	}

	return stream.CloseAndRecv()
}

func DiscoverAndDialGRPC() (*mdns.ServiceEntry, error) {
    entriesCh := make(chan *mdns.ServiceEntry, 3)

    go func() {
        err := mdns.Lookup(common.SERVICE_NAME, entriesCh)

        if err != nil {
            log.Println("❌ mDNS 查找失败:", err)
        }
    }()

    var entry *mdns.ServiceEntry

    select {
		case entry = <-entriesCh:
			log.Printf("🎯 发现服务: %s:%d\n", entry.AddrV4, entry.Port)
			return entry, nil
		case <-time.After(3 * time.Second):
			return nil, nil
	}
}

func main() {

	filePath := os.Args[1]

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	entry, err := DiscoverAndDialGRPC()

	if err != nil {
		log.Fatalln(err)
	}
	
	address := fmt.Sprintf("%s:%d", entry.AddrV4, entry.Port)

	// address = "127.0.0.1:60011"

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	client := pb.NewFileServiceClient(conn)

	err = UploadFile(client, filePath)

	if err != nil {
		log.Fatalln(err)
	}
}
