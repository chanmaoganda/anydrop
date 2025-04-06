#!/bin/bash
linux=linux-amd64
windows=windows-x86_64

out=../release

mkdir -p ${out}/${linux}
mkdir -p ${out}/${windows}

rm -r ${out}/${linux}/*
rm -r ${out}/${windows}/*

GOOS=linux
GOARCH=amd64

go build -o ${out}/${linux}/client ../client/client.go 
go build -o ${out}/${linux}/server ../server/server.go

GOOS=windows
GOARCH=amd64

go build -o ${out}/${windows}/client.exe ../client/client.go
go build -o ${out}/${windows}/server.exe ../server/server.go

cd $out

zip -r ${windows}.zip ${windows}
tar -czf ${linux}.tar.gz ${linux}
