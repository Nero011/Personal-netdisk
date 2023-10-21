package main

import (
	service "RpcProvider/kitex_gen/service"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// ProviderImpl implements the last service interface defined in the IDL.
type ProviderImpl struct{}

// User for user.json
type User struct {
	UserName string `json:"user"`
	Password string `json:"password"`
}

// CheckUser Check user.json
func CheckUser(username string, password string) (exist bool, ok bool) {
	file, err := os.Open("./config/user.json")
	if err != nil {
		fmt.Println("Open user.json err")
		return false, false
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var config []User
	user := make(map[string]string)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("json decode err")
		return false, false
	}
	for _, conf := range config {
		user[conf.UserName] = conf.Password
	}
	pwd, ok := user[username]
	if ok {
		if pwd == password {
			return true, true
		} else {
			return true, false
		}
	}
	return false, false
}

// Login implements the ProviderImpl interface.
func (s *ProviderImpl) Login(ctx context.Context, req *service.LoginRequest) (resp *service.LoginResponse, err error) {
	// TODO: Your code here...
	nameOk, pwdOk := CheckUser(req.UserName, req.UserPwd)
	if nameOk && pwdOk {
		resp = &service.LoginResponse{
			Success: true,
			ErrMsg:  "",
		}
	} else if nameOk && !pwdOk {
		resp = &service.LoginResponse{
			Success: false,
			ErrMsg:  "password error",
		}
	} else if !nameOk {
		resp = &service.LoginResponse{
			Success: false,
			ErrMsg:  "user not found",
		}
	}
	return
}

// Search implements the ProviderImpl interface.
func (s *ProviderImpl) Search(ctx context.Context, req *service.SearchRequest) (resp *service.SearchResponse, err error) {
	// TODO: Your code here...
	return
}

// Upload implements the ProviderImpl interface.
func (s *ProviderImpl) Upload(ctx context.Context, req *service.UploadResquest) (resp *service.UploadResponse, err error) {
	// TODO: Your code here...
	return
}

// Download implements the ProviderImpl interface.
func (s *ProviderImpl) Download(ctx context.Context, req *service.DownloadResquest) (resp *service.DownloadResponse, err error) {
	// TODO: Your code here...
	return
}

// Deletee implements the ProviderImpl interface.
func (s *ProviderImpl) Deletee(ctx context.Context, req *service.DeleteResquest) (resp *service.DeleteResponse, err error) {
	// TODO: Your code here...
	return
}

// Register implements the ProviderImpl interface.
func (s *ProviderImpl) Register(ctx context.Context, req *service.RegisterRequest) (resp *service.RegisterResponse, err error) {
	// TODO: Your code here...
	userExist, _ := CheckUser(req.UserName, req.UserPwd)
	if userExist {
		resp = &service.RegisterResponse{
			Success: false,
			ErrMsg:  "user is exist",
		}
		return nil, err
	}
	newUser := User{
		UserName: req.UserName,
		Password: req.UserPwd,
	}
	file, err := os.OpenFile("./config/user.json", os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()

	data, err := json.Marshal(newUser)
	if err != nil {
		fmt.Println("json marshal error")
		return nil, err
	}
	_, err = file.Seek(-1, io.SeekEnd)
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("write json error")
		return nil, err
	}
	resp = &service.RegisterResponse{
		Success: true,
		ErrMsg:  "",
	}
	return
}
