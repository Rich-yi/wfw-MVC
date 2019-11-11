package controller

import (
	"github.com/gin-gonic/gin"

	"fmt"
	getArea "wfw-MVC/proto/getArea"
	"net/http"
	"github.com/micro/go-micro/client"

	"context"

)
//获取所有地区信息
func GetArea(cxt *gin.Context){
	/*resp:=make(map[string]interface{})
	defer  cxt.JSON(http.StatusOK,resp)
	areas,err:=model.GetAreas()
	if err != nil {
		fmt.Println("获取所有地域信息失败")
		resp["errno"]=utils.RECODE_DBERR
		resp["errmsg"]=utils.RecodeText(utils.RECODE_DBERR)

		return
	}
	//把数据返回给前端
	resp["errno"]=utils.RECODE_OK
	resp["errmsg"]=utils.RecodeText(utils.RECODE_OK)
	resp["data"]=areas*/
	//调用远程服务,获取所有地域信息
	//初始化客户端
	microClient := getArea.NewGetAreaService("go.micro.srv.getArea",client.DefaultClient)
	//调用远程服务
	resp,err := microClient.MicroGetArea(context.TODO(),&getArea.Request{})
	if err != nil {
		fmt.Println("是这里打印的")
		fmt.Println(err)
		/*ctx.JSON(http.StatusOK,resp)
		return */
	}

	cxt.JSON(http.StatusOK,resp)

}