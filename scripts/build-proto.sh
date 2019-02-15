#!/bin/bash

export GO111MODULE=off
go get -u github.com/gogo/protobuf/proto
go get -u github.com/gogo/protobuf/protoc-gen-gogotypes
go get -u github.com/gogo/protobuf/gogoproto
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/mjibson/esc

echo "# Generating RPC Service..."
protoc -I=. \
  -I=$GOPATH/src \
  -I=$GOPATH/src/github.com/gogo/protobuf/protobuf \
  -I=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
  --gogotypes_out=\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
plugins=grpc,\
paths=source_relative:. \
./*.proto

ROOT=$(pwd)
echo "# Generating Plugins..."
# Generate various test protos.
PROTO_DIRS=(
  task
  plugins
)
for dir in ${PROTO_DIRS[@]}; do
  for p in `find $dir -name "*.proto"`; do
    if [[ $p == */import_public/* && ! $supportTypeAliases ]]; then
      echo "# $p (skipped)"
      continue;
    fi
    echo "## $p"
    protoc -I=$dir -I=$GOPATH/src --gogotypes_out=plugins=grpc,paths=source_relative:$dir \
      -I$ROOT \
      -I$GOPATH/src/github.com/gogo/protobuf/protobuf \
      -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
      -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
      -I$GOPATH/src/github.com/golang/protobuf/ptypes \
      $p
  done
done
