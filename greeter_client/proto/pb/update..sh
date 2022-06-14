#!/bin/bash

protoc --go_out=plugins=grpc:. --go_opt=Mgrpc.proto=../pb grpc.proto

### sed diff in mac os
sed -i '' 's/,omitempty//g' grpc.pb.go