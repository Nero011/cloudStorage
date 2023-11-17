package model

type LoginRequest struct {
	Name     string `json:"username" form:"username" api.raw:"username" query:"username"`
	Password string `json:"password" form:"password" api.raw:"password" query:"password"`
	Auth     string `json:"auth" form:"auth" api.raw:"auth" query:"auth"`
}

type LoginResponse struct {
	BasicResp BasicResponse `json:"basicResp"`
	Auth      string        `json:"auth"`
}

type RegisterRequest struct {
	Name     string `json:"username" form:"username" api.raw:"username" query:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	BasicResp BasicResponse `json:"basicResp"`
}
