package controller

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	pb "github.com/Urotea/go-grpc-sample/api"
	models "github.com/Urotea/go-grpc-sample/app/db/generated"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	pb.UnimplementedUserManagementServer
	logger *zap.SugaredLogger
	db     *sql.DB
}

var (
	DbError = errors.New("DBと接続でエラーが発生しました")
)

func RegisterUserManagementServer(logger *zap.SugaredLogger, db *sql.DB, server *grpc.Server) {
	pb.RegisterUserManagementServer(server, &UserServer{
		logger: logger,
		db:     db,
	})
}

func (server *UserServer) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserReply, error) {
	server.logger.Debugf("AddUserが呼ばれました。 req = %v", req)
	user := models.User{
		UserName:  fmt.Sprintf("%s %s", req.FirstName, req.LastName),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tx, err := server.db.BeginTx(ctx, nil)
	if err != nil {
		server.logger.Errorf("DBのトランザクションを開始しようとして失敗しました。error = %w", err)
		return nil, status.Errorf(codes.Internal, "書き込みに失敗しました。error = %w", err)
	}
	defer tx.Rollback()
	user.Insert(ctx, tx, boil.Infer())
	err = tx.Commit()
	if err != nil {
		server.logger.Errorf("コミットに失敗しました。error = %w", err)
		return nil, status.Errorf(codes.Internal, "書き込みに失敗しました。error = %w", err)
	}

	return &pb.AddUserReply{Message: "Hello " + req.FirstName + " " + req.LastName}, nil
}
