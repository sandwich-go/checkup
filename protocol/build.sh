#!/bin/bash
protoc --go_out=. ./protos/internal_command/internal.proto
protoc --go_out=. ./protos/common/common.proto
mv ./github.com/sandwich-go/internalcmd/protocol/internal_command/internal.pb.go ./gen/golang/internal_command/internal.pb.go
mv ./github.com/sandwich-go/internalcmd/protocol/common/common.pb.go ./gen/golang/common/common.pb.go
rm -rf ./github.com