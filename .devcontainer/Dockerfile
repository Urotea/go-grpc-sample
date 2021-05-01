FROM mcr.microsoft.com/vscode/devcontainers/go:0-1.16

ENV GO111MODULE=on
WORKDIR /workspace
RUN apt update && apt install -y protobuf-compiler
RUN go get google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
RUN cd /usr/bin && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | \
    tar xvz && \
    mv ./migrate.linux-amd64 migrate
RUN go install github.com/volatiletech/sqlboiler/v4@latest && \
    go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
ENV XDG_CONFIG_HOME=/workspace/go-grpc-sample/app/db
