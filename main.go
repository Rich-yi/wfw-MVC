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
	//数据库处理
	model.InitRedis()

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
	r1 := routher.Group("/api/v1.0")
	{
		r1.GET("/areas",controller.GetArea)
		r1.GET("/session",controller.GetSession)
		//传参方法,url传值,form表单传值,ajax传值,路径传值
		r1.GET("/imagecode/:uuid",controller.GetImageCd)
		r1.GET("/smscode/:mobile",controller.GetSmscd)
		r1.POST("/users",controller.PostRet)
	}
	routher.Run(":8001")
}
