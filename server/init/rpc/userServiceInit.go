package rpc

import "github.com/Nero011/cloudStorage/server/shared/kitex_gen/user/userservice"
import "github.com/Nero011/cloudStorage/server/global"

func InitUserService() {
	// TODO:init service discovery, and finish client option

	c, err := userservice.NewClient("userService")
	if err != nil {
	}

	global.GlobalUserClient = c
}
