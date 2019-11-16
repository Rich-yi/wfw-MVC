package main

import (
	"github.com/gin-gonic/gin"

	"wfw-MVC/model"
	"fmt"
	"wfw-MVC/controller"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-contrib/sessions"
	"wfw-MVC/utils"
	"net/http"

)

//路由过滤器
func Filter(ctx *gin.Context) {
	//登录校验
	session := sessions.Default(ctx)
	userName := session.Get("userName")
	resp := make(map[string]interface{})
	if userName == nil {
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
		ctx.JSON(http.StatusOK, resp)
		ctx.Abort()
		return
	}

	fmt.Println("next之前打印")

	//执行函数
	ctx.Next()

	fmt.Println("next之后打印....")
}
func MiddleTest() (func(*gin.Context)) {
	return func(context *gin.Context) {

	}
}
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
	routher.Static("/home", "view")
	//	routher.Use()
	store, err := redis.NewStore(20, "tcp", "127.0.0.1:6379", "", []byte("session"))
	if err != nil {
		fmt.Println("初始化session容器失败")
		return
	}
	/*store.Options(
		sessions.Options{
			MaxAge: 0,
		},
	)*/
	//路由使用中间件 gin中的session默认是生效时间是一个月
	/*routher.Use(sessions.Sessions("mySession", store))
	//使用路由的时候就可以使用session中间件了
	routher.GET("/session", func(context *gin.Context) {
		//初始化session对象
		se := sessions.Default(context)
		//设置session的时候除了set函数之外,必须调用save
		se.Set("test","bj5q")
		se.Save()

		context.Writer.WriteString("设置session")
	})

	//获取session
	routher.GET("/getSession", func(context *gin.Context) {
		//初始化session对象
		se := sessions.Default(context)
		//获取session
		result := se.Get("test")
		fmt.Println("得到的session数据为",result.(string))

		context.Writer.WriteString("获取session")
	})

	//测试
	routher.GET("/test", func(context *gin.Context) {
		//设置cookie  cookie有两种,一种是有时间效应的,一种是临时cookie
		*//*context.SetCookie("myTest","bj5q",0,"","",false,true)
		context.Writer.WriteString("测试cookie")*//*
	//请求分配*/
	r1 := routher.Group("/api/v1.0")
	{
		r1.GET("/areas", controller.GetArea)
		//传参方法,url传值,form表单传值,ajax传值,路径传值
		r1.GET("/imagecode/:uuid", controller.GetImageCd)
		r1.GET("/smscode/:mobile", controller.GetSmscd)
		r1.POST("/users", controller.PostRet)
		//登录业务
		r1.Use(sessions.Sessions("mysession", store))
		r1.POST("/sessions", controller.PostLogin)
		r1.GET("/session", controller.GetSession)
		//r1.Use(MiddleTest())
		r1.Use(Filter)
		r1.DELETE("/session", controller.DeleteSession)
		r1.GET("/user",controller.GetUserInfo)
		r1.PUT("/user/name",controller.PutUserInfo)

		r1.POST("/user/avatar",controller.PostAvatar)
	}
	routher.Run(":8001")
}
