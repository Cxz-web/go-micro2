package service

import (
	"context"
	"github.com/Cxz-web/go-micro2/user/proto/user"
)

type UserService struct {
}

func (u *UserService) Register(ctx context.Context, req *user.UserRegisterRequest, response *user.UserRegisterResponse) error {
	return nil
}

func (u *UserService) Login(ctx context.Context, req *user.UserLoginRequest, response *user.UserLoginResponse) error {
	return nil
}

func (u *UserService) GetUserInfo(ctx context.Context, req *user.UserInfoRequest, response *user.UserInfoResponse) error {
	return nil
}

//
//rpc Register(UserRegisterRequest) returns (UserRegisterResponse) {}
//rpc Login(UserLoginRequest) returns (UserLoginResponse) {}
//rpc GetUserInfo(UserInfoRequest) returns (UserInfoResponse) {}
