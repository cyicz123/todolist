package service

import (
	"context"

	"github.com/cyicz123/todolist/user/internal/handler"
	"github.com/cyicz123/todolist/user/pkg/e"
	"github.com/cyicz123/todolist/user/pkg/repo"
)

type UserService struct {
	UnimplementedUserServiceServer
	user			handler.UserInterface
}

func NewUserService(user handler.UserInterface) *UserService {
	return &UserService{user: user}
}

func (u *UserService) UserLogin(ctx context.Context,req *UserRequest) (*UserDetailResponse, error) {
	info, _ := u.req2userModel(req)
	if err := u.user.Login(info); err != nil {
		return &UserDetailResponse{Code: uint32(err.Code())}, err
	}
	return &UserDetailResponse{Code: uint32(e.SUCCESS.Code())}, nil
}


func (u *UserService) UserRegister(ctx context.Context, req *UserRequest) (*UserDetailResponse, error) {
	info, _ := u.req2userModel(req)
	err := u.user.Register(info)
	if err != nil {
		return &UserDetailResponse{Code: uint32(err.Code())}, err
	}
	return &UserDetailResponse{Code: uint32(e.SUCCESS.Code())}, nil
}

func (u *UserService) UserLogout(ctx context.Context, req *UserRequest) (*UserDetailResponse, error) {
	return &UserDetailResponse{Code: uint32(e.SUCCESS.Code())}, nil
}

func (u *UserService) UserDelete(ctx context.Context, req *UserRequest) (*UserDetailResponse, error) {
	info, _ := u.req2userModel(req)
	err := u.user.UserDelete(info.UserName)
	if err != nil {
		return &UserDetailResponse{Code: uint32(err.Code())}, err
	}
	return &UserDetailResponse{Code: uint32(e.SUCCESS.Code())}, nil
}

func (u *UserService) req2userModel(req *UserRequest) (*repo.User, e.ErrInterface) {
	user := &repo.User{}
	user.UserName = req.UserName
	user.NickName = req.NickName
	user.PasswordDigest = req.Password
	return user, nil
}