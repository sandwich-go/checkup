#!/bin/bash
protoc --go_out=. ./protos/internal_command/internal.proto
mv ./github.com/sandwich-go/internalcmd/protocol/internal_command/internal.pb.go ./gen/golang/internal_command/internal.pb.go
rm -rf ./github.com