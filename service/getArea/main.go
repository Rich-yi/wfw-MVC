package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"wfw-MVC/service/getArea/handler"
	//"wfw-MVC/service/getArea/subscriber"

	getArea "wfw-MVC/service/getArea/proto/getArea"
	"wfw-MVC/service/getArea/model"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.getArea"),
		micro.Version("latest"),
	)
	model.InitDb()
	model.InitRedis()
	// Initialise service
	service.Init()

	// Register Handler
	getArea.RegisterGetAreaHandler(service.Server(), new(handler.GetArea))

	//// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.getArea", service.Server(), new(subscriber.GetArea))
	//
	//// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.getArea", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
