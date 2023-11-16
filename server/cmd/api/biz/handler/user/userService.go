package user

import (
	"context"
	"github.com/Nero011/cloudStorage/server/shared/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func Login(ctx *context.Context, c *app.RequestContext) {
	var req user.LoginRequest      //FIXME: this struct is not be define yet, define it on api/model, is not kitex.requset
	err := c.BindAndValidate(&req) // bind http request to struct

	if err != nil {
		c.JSON(http.StatusBadRequest, "") //TODO:packet responce
	}

	// rpc call
	resp := global.UserClient.Login(ctx, &user.LoginRequest{
		UserName:     req.UserName,
		UserPassword: req.UserPassword,
	})

	if !resp.Success {
		c.JSON(http.StatusBadRequest, "rpc error")
	}
	c.JSON(http.StatusOK, "todo:request with auth")

}
