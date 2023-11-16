package router

import "github.com/cloudwego/hertz/pkg/app/server"
import user "github.com/Nero011/cloudStorage/server/cmd/api/biz/handler/user"

func Register(r *server.Hertz) {
	root := r.Group("/")
	{
		login := root.Group("/login")
		login.POST("/user", user.Login)
	}
}
