namespace go service

struct RegisterRequest{
    1: string UserName
    2: string UserPwd
}
struct RegisterResponse{
    1: bool Success
    2: string ErrMsg
}

struct LoginRequest{
    1: string UserName
    2: string UserPwd
}

struct LoginResponse{
    1: bool Success
    2: string ErrMsg
}

struct SearchRequest{
    1: string FileName
}

struct SearchResponse{
    1: bool Success
    2: string FilePath
}

struct UploadResquest{
    1: string FilePath //name + path
    2: i64 FileSize //kb
}
struct UploadResponse{
    1: bool Success
    2: string StoreAddr

}

struct DownloadResquest{
    1: string FilePath
}
struct DownloadResponse{
    1: bool Success
    2: string StoreAddr
    3: i64 FileSize
}

struct DeleteResquest{
    1: string FilePath
}
struct DeleteResponse{
    1: bool Success
}

service Provider{
    RegisterResponse Register(1:RegisterRequest req)
    LoginResponse Login(1:LoginRequest req)
    SearchResponse Search(1:SearchRequest req)
    UploadResponse Upload(1:UploadResquest req)
    DownloadResponse Download(1:DownloadResquest req)
    DeleteResponse  Deletee(1:DeleteResquest req)

}