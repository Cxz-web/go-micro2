package  main

// protoc --proto_path=./ --micro_out=./ --go_out=./ ./proto/*.proto
import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"new_go_study/src/proto"
)

func main() {
	service := micro.NewService(micro.Name("cxz.client"))
	service.Init()
	// 注册服务
	serviceRemote := proto.NewClientAService("cxz.server", service.Client())

	res, err := serviceRemote.SayHello(context.TODO(), &proto.SayRequest{Messge: "测试的"})
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
