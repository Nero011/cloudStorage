package main

import (
	user "github.com/Nero011/cloudStorage/server/shared/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:12304")
	svr := user.NewServer(new(UserServiecImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
