package rpc

import (
	"github.com/Nero011/cloudStorage/server/shared/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
)
import "github.com/Nero011/cloudStorage/server/global"

func InitUserService() {
	// TODO:init service discovery, and finish client option

	c, err := userservice.NewClient("userService", client.WithHostPorts("0.0.0.0:12304"))
	if err != nil {
	}

	global.GlobalUserClient = c
}
