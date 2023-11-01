package handler

import (
	service "RpcProvider/kitex_gen/service"
	mysqlutil "RpcProvider/util/mysql"
	"context"
	"fmt"
)

// ProviderImpl implements the last service interface defined in the IDL.
type ProviderImpl struct{}

// Register implements the ProviderImpl interface.
func (s *ProviderImpl) Register(ctx context.Context, req *service.RegisterRequest) (resp *service.RegisterResponse, err error) {
	// TODO: Your code here...
	db := mysqlutil.Init()
	//sqlStr := "select user_name from user where	user_name = \"" + req.GetUserName() + "\""
	sqlStr := fmt.Sprintf(`select user_name from user where user_name = "%s"`, req.GetUserName())

	rows, err := db.Query(sqlStr)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	// 检测是否数据库中已存在用户
	exist := false
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		exist = true
		if err != nil {
			return nil, err
		}
	}
	if exist {
		resp = &service.RegisterResponse{
			Success: false,
			ErrMsg:  "user is exist",
		}
		return resp, nil
	}

	// 在数据库中插入新用户的数据
	sqlStr = fmt.Sprintf(`insert into user (user_name, user_pwd) values ("%s", "%s")`, req.GetUserName(), req.GetUserPwd())
	res, err := db.Exec(sqlStr)
	if err != nil {
		return nil, err
	}
	aff, _ := res.RowsAffected()
	println(aff)
	resp = &service.RegisterResponse{
		Success: true,
		ErrMsg:  "",
	}
	return resp, nil
}

// Login implements the ProviderImpl interface.
func (s *ProviderImpl) Login(ctx context.Context, req *service.LoginRequest) (resp *service.LoginResponse, err error) {
	// TODO: Your code here...
	db := mysqlutil.Init()
	sqlStr := fmt.Sprintf(`select user_name, user_pwd from user where user_name = "%s"`, req.GetUserName())
	row, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	exist := false
	for row.Next() {
		exist = true
		var name, pwd string
		err = row.Scan(&name, &pwd)
		if err != nil {
			return nil, err
		}
		if pwd == req.GetUserPwd() {
			resp = &service.LoginResponse{
				Success: true,
				ErrMsg:  "",
			}
			return resp, nil
		} else {
			resp = &service.LoginResponse{
				Success: false,
				ErrMsg:  "password error",
			}
		}
	}
	if exist == false {
		resp = &service.LoginResponse{
			Success: false,
			ErrMsg:  "user is not exist",
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

// Delete implements the ProviderImpl interface.
func (s *ProviderImpl) Delete(ctx context.Context, req *service.DeleteResquest) (resp *service.DeleteResponse, err error) {
	// TODO: Your code here...
	return
}
