package main

import (
	"context"
	userservice "github.com/Nero011/cloudStorage/server/shared/kitex_gen/user"
	"github.com/Nero011/cloudStorage/server/shared/model/user"
	"github.com/Nero011/cloudStorage/tools/mysql"
	"github.com/cloudwego/kitex/pkg/klog"
)

// UserServiecImpl implements the last service interface defined in the IDL.
type UserServiecImpl struct{}

// Register implements the UserServiecImpl interface.
func (s *UserServiecImpl) Register(ctx context.Context, req *userservice.RegisterRequest) (resp *userservice.RegisterResponse, err error) {
	u := user.User{
		Name:     req.GetUserName(),
		Password: req.GetUserPassword(),
	}
	if mysql.CheckUser(&u) {
		resp = &userservice.RegisterResponse{
			Success: false,
			ErrMsg:  "user is exist",
		}
	}
	err = mysql.CreateUser(&u)
	if err != nil {
		klog.Error("register error:", err)
		resp = &userservice.RegisterResponse{
			Success: false,
			ErrMsg:  "register service error",
		}
		return nil, err
	}
	resp = &userservice.RegisterResponse{
		Success: true,
		ErrMsg:  "",
	}
	return
}

// Login implements the UserServiecImpl interface.
func (s *UserServiecImpl) Login(ctx context.Context, req *userservice.LoginRequest) (resp *userservice.LoginResponse, err error) {
	// TODO: Your code here...
	return
}
