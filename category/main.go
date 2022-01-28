package main

import (
	"fmt"
	"github.com/Cxz-web/go-micro2/category/common"
	"github.com/Cxz-web/go-micro2/category/handler"
	pb "github.com/Cxz-web/go-micro2/category/proto"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	log "github.com/micro/micro/v3/service/logger"
	"github.com/opentracing/opentracing-go"
)

func main() {

	db, err := gorm.Open("mysql", "root:watch8562871@(127.0.0.1:3306)/category?charset=utf8mb4&parseTime=True&loc=Local")

	defer db.Close()

	// 禁止副表
	db.SingularTable(true)

	//db.CreateTable(&model.Category{})

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	// 注册中心
	consuleRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	t, io, err := common.NewTracer("go.micro.service.category", "127.0.0.1:6831")

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	defer io.Close()

	opentracing.SetGlobalTracer(t)

	service := micro.NewService(
		micro.Name("category"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8082"),
		micro.Registry(consuleRegistry),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	service.Init()

	// Register handler
	pb.RegisterCategoryHandler(service.Server(), handler.NewCategoryHandler())

	// Run service
	if err := service.Run(); err != nil {
		fmt.Println("micro run error")
		return
	}

	fmt.Println("successful")
}
