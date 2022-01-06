package  main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"new_go_study/src/proto"
)

type CxzServer struct {
}

func (c *CxzServer) SayHello(ctx context.Context,req *proto.SayRequest, response *proto.SayResponse,) error {
    response.Answer = "我们的口号是：" + req.Messge
	return nil
}

func main() {
	service := micro.NewService(micro.Name("cxz.server"))
	service.Init()
	// 注册服务
	proto.RegisterClientAHandler(service.Server(), new(CxzServer))
	// 运行服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}






}
