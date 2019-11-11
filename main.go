package main

import (
	"github.com/gin-gonic/gin"

	"wfw-MVC/model"
	"fmt"
	"wfw-MVC/controller"
)

func main() {
	//初始化路由
	routher := gin.New()
	err := model.InitDb()
	if err != nil {
		fmt.Println("数据库创建失败", err)
		return
	}
	//路由模块
	//router.Group()
	//展示静态页面
	//	静态路由
	routher.Static("/home","view")
	//请求分配
	r1 := routher.Group("/api/V1.0")
	{
		r1.GET("/areas",controller.GetArea)

	}
	routher.Run(":8001")
}
