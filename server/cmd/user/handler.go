package main

import (
	"context"
	userservice "github.com/Nero011/cloudStorage/server/shared/kitex_gen/userservice"
)

// UserServiecImpl implements the last service interface defined in the IDL.
type UserServiecImpl struct{}

// Register implements the UserServiecImpl interface.
func (s *UserServiecImpl) Register(ctx context.Context, req *userservice.RegisterRequest) (resp *userservice.RegisterResponse, err error) {

	return
}

// Login implements the UserServiecImpl interface.
func (s *UserServiecImpl) Login(ctx context.Context, req *userservice.LoginRequest) (resp *userservice.LoginResponse, err error) {
	// TODO: Your code here...
	return
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *userservice.RegisterRequest) (resp *userservice.RegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *userservice.LoginRequest) (resp *userservice.LoginResponse, err error) {
	// TODO: Your code here...
	return
}
