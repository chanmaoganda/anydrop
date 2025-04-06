#!/bin/bash
out=../bins

rm ${out}/client ${out}/server

go build -o ${out}/client ../cmd/client/client.go 
go build -o ${out}/server ../cmd/server/server.go
