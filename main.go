package main

//go:generate sqlboiler psql
//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/sample.proto

import (
	"database/sql"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/Urotea/go-grpc-sample/app/controller"
	"github.com/Urotea/go-grpc-sample/app/logger"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

var (
	port       = flag.String("port", getEnvDefault("PORT", "9000"), "サーバのポート番号を設定する")
	dbName     = flag.String("db_name", getEnvDefault("DB_NAME", ""), "データベース名を設定する")
	dbHost     = flag.String("db_host", getEnvDefault("DB_HOST", ""), "DBのHOSTを設定する。")
	dbPort     = flag.String("db_port", getEnvDefault("DB_PORT", ""), "DBのPORTを設定する。")
	dbUser     = flag.String("db_user", getEnvDefault("DB_USER", ""), "DBのユーザを設定する。")
	dbPassword = flag.String("db_password", getEnvDefault("DB_PASSWORD", ""), "DBのパスワードを設定する。")
)

func main() {
	flag.Parse()
	logger := logger.GetLogger()
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		*dbHost, *dbPort, *dbUser, *dbPassword, *dbName))
	if err != nil {
		logger.Fatalf("DBへの接続に失敗しました。 err = %w", err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	logger.Debugf("UserManagementServerを登録します。")
	controller.RegisterUserManagementServer(logger, db, s)

	logger.Debugf("サーバが起動します。")
	if err := s.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}

func getEnvDefault(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}
