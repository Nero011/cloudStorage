package main

import (
	user "github.com/Nero011/cloudStorage/server/shared/kitex_gen/user/userservice"
	"log"
)

func main() {
	svr := user.NewServer(new(UserServiecImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
