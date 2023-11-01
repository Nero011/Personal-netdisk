package handler

import (
	"RpcProvider/kitex_gen/service"
	"context"
	"reflect"
	"testing"
)

func TestProviderImpl_Register(t *testing.T) {
	type args struct {
		ctx context.Context
		req *service.RegisterRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *service.RegisterResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{context.Background(), &service.RegisterRequest{
			UserName: "wasd",
			UserPwd:  "111",
		}}, wantResp: &service.RegisterResponse{
			Success: true,
			ErrMsg:  "",
		},
			wantErr: false,
		},
		{name: "test2", args: args{context.Background(), &service.RegisterRequest{
			UserName: "alu",
			UserPwd:  "111",
		}}, wantResp: &service.RegisterResponse{
			Success: false,
			ErrMsg:  "user is exist",
		},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ProviderImpl{}
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

func TestProviderImpl_Login(t *testing.T) {
	type args struct {
		ctx context.Context
		req *service.LoginRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *service.LoginResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
		{"user not exist", args{
			ctx: context.Background(),
			req: &service.LoginRequest{
				UserName: "admi",
				UserPwd:  "123456",
			},
		}, &service.LoginResponse{
			Success: false,
			ErrMsg:  "user is not exist",
		}, false},
		{"pwd error", args{
			ctx: context.Background(),
			req: &service.LoginRequest{
				UserName: "admin",
				UserPwd:  "1234",
			},
		}, &service.LoginResponse{
			Success: false,
			ErrMsg:  "password error",
		}, false},
		{"success", args{
			ctx: context.Background(),
			req: &service.LoginRequest{
				UserName: "admin",
				UserPwd:  "123456",
			},
		}, &service.LoginResponse{
			Success: true,
			ErrMsg:  "",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ProviderImpl{}
			gotResp, err := s.Login(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Login() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
