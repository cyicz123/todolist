package main

import (
	"net"

	"github.com/cyicz123/todolist/user/config"
	"github.com/cyicz123/todolist/user/internal/handler"
	"github.com/cyicz123/todolist/user/internal/service"
	"github.com/cyicz123/todolist/user/pkg/logger"
	"github.com/cyicz123/todolist/user/pkg/repo"

	"google.golang.org/grpc"
)

func main() {
	l := logger.New("user")
	v := config.GetInstance()
	dbFactory := &repo.MysqlFactory{}
	r, err := dbFactory.New(l, v)
	if err != nil {
		l.Panic(err)
	}
	u := handler.NewUserInfo(r, l)

	server := grpc.NewServer()
	defer server.Stop()

	grpcAddress := v.GetString("server.grpcAddress")
	service.RegisterUserServiceServer(server, service.NewUserService(u))
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		l.Panic(err)
	}
	l.Info("Server started listen on ", grpcAddress)
	if err := server.Serve(lis); err != nil {
		l.Panic(err)
	}
}