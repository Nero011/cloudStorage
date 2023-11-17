package user

import (
	"context"
	huser "github.com/Nero011/cloudStorage/server/cmd/api/biz/model"
	"github.com/Nero011/cloudStorage/server/global"
	kuser "github.com/Nero011/cloudStorage/server/shared/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"net/http"
)

func Login(ctx context.Context, c *app.RequestContext) {
	var req huser.LoginRequest
	var resp *huser.LoginResponse
	err := c.BindJSON(&req) // bind http request to struct

	if err != nil {
		resp = &huser.LoginResponse{
			BasicResp: huser.BasicResponse{http.StatusBadRequest, "Err request"},
			Auth:      "",
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// rpc call
	uresp, err := global.GlobalUserClient.Login(ctx, &kuser.LoginRequest{
		UserName:     req.Name,
		UserPassword: req.Password,
	})
	if err != nil {
		hlog.CtxErrorf(ctx, "user rpc call error ?", err)
	}

	if !uresp.Success {
		resp = &huser.LoginResponse{
			BasicResp: huser.BasicResponse{http.StatusBadRequest, "bad request"},
			Auth:      "",
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp = &huser.LoginResponse{
		BasicResp: huser.BasicResponse{http.StatusOK, "ok"},
		Auth:      "test auth",
	}
	c.JSON(http.StatusOK, resp)

}

func Register(ctx context.Context, c *app.RequestContext) {
	var req huser.RegisterRequest
	var resp *huser.RegisterResponse
	err := c.BindJSON(&req)
	if err != nil {
		resp = &huser.RegisterResponse{
			BasicResp: huser.BasicResponse{http.StatusBadRequest, "request bind wrong"},
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	uresp, err := global.GlobalUserClient.Register(ctx, &kuser.RegisterRequest{
		UserName:     req.Name,
		UserPassword: req.Password,
	})
	if err != nil {
		resp = &huser.RegisterResponse{
			BasicResp: huser.BasicResponse{
				StatusCode: http.StatusInternalServerError,
				StatusMsg:  "user rpc call error",
			},
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	if !uresp.Success {
		resp = &huser.RegisterResponse{BasicResp: huser.BasicResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  uresp.ErrMsg,
		}}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp = &huser.RegisterResponse{BasicResp: huser.BasicResponse{
		StatusCode: http.StatusOK,
		StatusMsg:  "",
	}}
	c.JSON(http.StatusOK, resp)
	return
}
