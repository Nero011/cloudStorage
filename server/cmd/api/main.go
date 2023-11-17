// Code generated by hertz generator.

package main

import (
	"github.com/Nero011/cloudStorage/server/cmd/api/biz/router"
	"github.com/Nero011/cloudStorage/server/init/rpc"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default()

	router.Register(h)
	rpc.InitUserService()

	register(h)
	h.Spin()
}
