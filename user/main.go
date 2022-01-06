package main

import (
	"fmt"
	userSrv "github.com/Cxz-web/go-micro2/user/domain/service"
	user "github.com/Cxz-web/go-micro2/user/proto/user"
	"github.com/micro/go-micro/v2"
)

func main() {
	// Create service
	service := micro.NewService(
		micro.Name("user"),
		micro.Version("latest"),
	)

	service.Init()

	// Register handler
	user.RegisterUserHandler(service.Server(), new(userSrv.UserService))

	// Run service
	if err := service.Run(); err != nil {
		fmt.Println("micro run error")
	}

	fmt.Println("successful")
}
