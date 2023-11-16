namespace go user

struct BasicResponse{
    1: bool     success
    2: string   errMsg
}

struct LoginRequest{
    1: string   name (api.raw = "name")
    2: string   password (api.raw = "password")
}

struct LoginResponse{
    1: BasicResponse    bresp
    2: string           auth
}

service UserService{
    LoginResponse Login(1: LoginRequest req)
}