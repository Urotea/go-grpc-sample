FROM mcr.microsoft.com/vscode/devcontainers/go:0-1.16

ENV GO111MODULE=on
RUN apt update && apt install -y protobuf-compiler
RUN go get google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
