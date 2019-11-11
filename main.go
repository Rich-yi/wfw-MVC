package main

import (
	"github.com/gin-gonic/gin"

	"wfw-MVC/model"
)

func main() {
	//初始化路由
	routher := gin.New()

	//请求分配
	r1 := routher.Group("V1")
	{
		r1.GET("/cat", func(context *gin.Context) {
			context.Writer.WriteString("bigcat")
		})


	}
	model.InitModel()
	model.InitData()
	model.SearchData()
	routher.Run(":8001")
}
