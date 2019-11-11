package controller

import (
	"github.com/gin-gonic/gin"
	"wfw-MVC/model"
	"fmt"
	"wfw-MVC/utils"
	"net/http"
)
//获取所有地区信息
func GetArea(cxt *gin.Context){
	resp:=make(map[string]interface{})
	defer  cxt.JSON(http.StatusOK,resp)
	areas,err:=model.GetAreas()
	if err != nil {
		fmt.Println("获取所有地域信息失败")
		resp["errno"]=utils.RECODE_DBERR
		resp["errmsg"]=utils.RecodeText(utils.RECODE_DBERR)

		return
	}
	//把数据返回给前端
	resp["errno"]=utils.RECODE_DBERR
	resp["errmsg"]=utils.RecodeText(utils.RECODE_DBERR)
	resp["data"]=areas

}