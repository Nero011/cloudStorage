package main

import (
	"context"
	user "github.com/Nero011/cloudStorage/server/shared/kitex_gen/user"
)

// UserServiecImpl implements the last service interface defined in the IDL.
type UserServiecImpl struct{}

// Register implements the UserServiecImpl interface.
func (s *UserServiecImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// Login implements the UserServiecImpl interface.
func (s *UserServiecImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	// TODO: Your code here...
	return
}
