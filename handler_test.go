package main

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
	req := service.RegisterRequest{
		UserName: "elo",
		UserPwd:  "111",
	}
	resp := service.RegisterResponse{
		Success: true,
		ErrMsg:  "",
	}
	tests := []struct {
		name     string
		args     args
		wantResp *service.RegisterResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
		{"version1", args{context.Background(), &req}, &resp, false},
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
