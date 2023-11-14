package main

import (
	"context"
	userservice "github.com/Nero011/cloudStorage/server/shared/kitex_gen/user"
	"reflect"
	"testing"
)

func TestUserServiecImpl_Register(t *testing.T) {
	type args struct {
		ctx context.Context
		req *userservice.RegisterRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *userservice.RegisterResponse
		wantErr  bool
	}{
		{"test1",
			args{
				ctx: context.Background(),
				req: &userservice.RegisterRequest{
					UserName:     "adls",
					UserPassword: "1223",
				},
			},
			&userservice.RegisterResponse{
				Success: true,
				ErrMsg:  "",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserServiecImpl{}
			gotResp, err := s.Register(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Register() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
