package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"wfw-MVC/service/getImage/handler"


	getImage "wfw-MVC/service/getImage/proto/getImage"
	"github.com/micro/go-micro/registry/consul"
	"wfw-MVC/service/getImage/model"
)

func main() {
	//使用consul做服务发现
	consulReg:=consul.NewRegistry()
	model.InitRedis()
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.getImage"),
		micro.Version("latest"),
		micro.Registry(consulReg),
		//隐藏bug   注册的服务注销了吗   不一定注销 65535
		micro.Address(":9981"),

	)

	// Initialise service
	service.Init()

	// Register Handler
	getImage.RegisterGetImageHandler(service.Server(), new(handler.GetImage))


	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
