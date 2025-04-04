protos:
	protoc --go_out=../filetransfer --go_opt=paths=source_relative \
    --go-grpc_out=../filetransfer --go-grpc_opt=paths=source_relative \
	file.proto

.PHONY: clean
clean:
	rm -rf ../filetransfer/*