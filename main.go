package main

import (
	"context"
	"flag"
	"net"

	pb "github.com/Urotea/go-grpc-sample/api"
	"github.com/Urotea/go-grpc-sample/logger"
	"github.com/Urotea/go-grpc-sample/postgres"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (server *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Debugf("SayHelloが呼ばれました。 req = %s", req.Name)
	postgresDao.Create(postgres.User{UserName: req.Name})
	return &pb.HelloReply{Message: "Hello " + req.Name}, nil
}

var (
	schema = flag.String("schema", "", "DBのDSNを設定する")
	log    = logger.GetLogger()
)

var postgresDao *postgres.PostgresDao = nil

func main() {
	flag.Parse()
	var err error
	postgresDao, err = postgres.New(*schema)
	if err != nil {
		logger.Fatalf("DBへの接続に失敗しました。 error = %s", err)
	}
	logger.Debugf("DBに接続しました。schema = %s", *schema)

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	logger.Debugf("サーバが起動しました。")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
