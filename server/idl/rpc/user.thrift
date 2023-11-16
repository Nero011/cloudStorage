namespace go userservice




struct RegisterRequest{
    1: string userName
    2: string userPassword
}

struct RegisterResponse{
    1: bool success
    2: string errMsg
}

struct LoginRequest {
    1: string userName
    2: string userPassword
}

struct LoginResponse{
    1: bool success
    2: string errMsg
}

service UserService{
    RegisterResponse Register(1:RegisterRequest req)
    LoginResponse Login(1:LoginRequest req)
}