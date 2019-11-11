package main

import (
	"github.com/gin-gonic/gin"

	"wfw-MVC/model"
	"fmt"
)

func main() {
	//初始化路由
	routher := gin.New()
	err:=model.InitDb()
	if err != nil {
		fmt.Println("数据库创建失败",err)
		return
	}
	//请求分配
	r1 := routher.Group("V1")
	{
		r1.GET("/cat", func(context *gin.Context) {
			context.Writer.WriteString("bigcat")
		})

	}
	routher.Run(":8001")
}
