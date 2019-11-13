package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"wfw-MVC/service/register/handler"

	register "wfw-MVC/service/register/proto/register"
	"github.com/micro/go-micro/registry/consul"
	"wfw-MVC/service/register/model"
)

func main() {
	RegisterConsul := consul.NewRegistry()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.register"),
		micro.Version("latest"),
		micro.Registry(RegisterConsul),
		micro.Address(":9998"),
	)

	// Initialise service
	service.Init()
	model.InitRedis()
	model.InitDb()
	// Register Handler
	register.RegisterRegisterHandler(service.Server(),new(handler.Register))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
