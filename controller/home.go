package controller

import (
	"github.com/gin-gonic/gin"

	"fmt"
	getArea "wfw-MVC/proto/getArea"
	"net/http"

	"context"

	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro"
	"wfw-MVC/utils"
	getImage "wfw-MVC/service/getImage/proto/getImage"
	"github.com/afocus/captcha"
	"encoding/json"
	"image/png"

	register "wfw-MVC/service/register/proto/register"
	"regexp"
)

//获取所有地区信息
func GetArea(cxt *gin.Context) {
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
	microtoConsul := consul.NewRegistry()
	microclient := micro.NewService(
		micro.Registry(microtoConsul),
	)
	microClient := getArea.NewGetAreaService("go.micro.srv.getArea", microclient.Client())
	//调用远程服务
	resp, err := microClient.MicroGetArea(context.TODO(), &getArea.Request{})
	if err != nil {
		fmt.Println("是这里打印的")
		fmt.Println(err)
		/*ctx.JSON(http.StatusOK,resp)
		return */
	}

	cxt.JSON(http.StatusOK, resp)

}

//写一个假的session请求
func GetSession(ctx *gin.Context) {
	//构造未登录
	resp := make(map[string]interface{})

	resp["errno"] = utils.RECODE_LOGINERR
	resp["errmsg"] = utils.RecodeText(utils.RECODE_LOGINERR)

	ctx.JSON(http.StatusOK, resp)

}

//获取验证码图片的方法
func GetImageCd(ctx *gin.Context) {
	//获取数据
	uuid := ctx.Param("uuid")
	//校验数据
	if uuid == "" {
		fmt.Println("获取数据错误")
		return
	}

	//处理数据
	//调用远程服务
	//初始化客户端
	consulReg:=consul.NewRegistry()
	microService:=micro.NewService(
		micro.Registry(consulReg),
	)
	microClient:=getImage.NewGetImageService("go.micro.srv.getImage",microService.Client())
	//调用远程服务
	resp,err:=microClient.MicroGetImg(context.TODO(),&getImage.Request{Uuid:uuid})

	//获取数据
	if err != nil {
		fmt.Println("获取远端数据失败")
		ctx.JSON(http.StatusOK,resp)
		return
	}
	//返回json数据
	//反序列化拿到img数据
	var img captcha.Image
	json.Unmarshal(resp.Data,&img)
	png.Encode(ctx.Writer,img)

}
//发送短信验证码
func GetSmscd(ctx *gin.Context)  {
	//获取数据
	mobile:=ctx.Param("mobile")

	//获取输入的图片验证码
	text := ctx.Query("text")
	//获取验证码图片的uuid
	uuid := ctx.Query("id")
	if mobile==""||text==""||uuid==""{
		fmt.Println("传入数据不完整" )
		return
	}
	//处理数据  放在服务端处理
	//初始化客户端
	microClient:=register.NewRegisterService("go.micro.srv.register",utils.GetMicroClient())
	//调用远程客户端
	resp,err:=microClient.SmsCode(context.TODO(),&register.Request{
		Uuid:uuid,
		Text:text,
		Mobile:mobile,
	})
	if err != nil {
		fmt.Println("调用远程服务错误",err)
		/*ctx.JSON(http.StatusOK,resp)
		return*/
	}

	ctx.JSON(http.StatusOK,resp)
}
//获取数据
type RegisterUser struct {
	Mobile string `json:"mobile"`
	Password string  `json:"password"`
	SmsCode string    `json:"sms_code"`
}
func PostRet(ctx *gin.Context)  {
	//确定容器
	resp := make(map[string]interface{})
	//绑定数据
	var regUser RegisterUser
	err := ctx.Bind(&regUser)
	if err != nil {
		resp["errno"] = utils.RECODE_DATAERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR)
		ctx.JSON(200,resp)
		return
	}
	//校验数据
	if regUser.Mobile == "" || regUser.Password == "" || regUser.SmsCode == ""{
		resp["errno"] = utils.RECODE_DATAERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR)
		ctx.JSON(200,resp)
		return
	}
	reg,_:=regexp.Compile(`^1[3,4,5,7,8]\d{9}$`)
	if !reg.MatchString(regUser.Mobile){
		resp["errno"] = utils.RECODE_DATAERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR)
		ctx.JSON(200,resp)
		return
	}
	microClient:=register.NewRegisterService("go.micro.srv.register",utils.GetMicroClient())
	regResp,err:=microClient.Register(context.TODO(),&register.RegRequest{

		Mobile:regUser.Mobile,
		Password:regUser.Password,
		SmsCode:regUser.SmsCode,
	})
	if err != nil {
		fmt.Println("获取远端数据失败")
		return
	}

//返回数据
ctx.JSON(http.StatusOK,&regResp)

}
