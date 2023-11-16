package model

type LoginRequest struct {
	Name     string `json:"username" form:"username" api.raw:"username" query:"username"`
	Password string `json:"password" form:"password" api.raw:"password" query:"password"`
	Auth     string `json:"auth" form:"auth" api.raw:"auth" query:"auth"`
}

type LoginResponse struct {
	Bresp BasicResponse `json:"bresp"`
	Auth  string        `json:"auth"`
}
