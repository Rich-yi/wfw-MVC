package handler

import (
	"context"
	register "wfw-MVC/service/register/proto/register"
	"wfw-MVC/service/register/model"
	"wfw-MVC/service/register/utils"
	"errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"math/rand"
	"time"
	"fmt"

	"crypto/sha256"
)

type Register struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Register) SmsCode(ctx context.Context, req *register.Request, rsp *register.Response) error {
	//写具体业务   uuid   text    mobile

	//验证图片验证码是否输入正确
	rnd, err := model.GetImgCode(req.Uuid)
	if err != nil {
		rsp.Errno = utils.RECODE_NODATA
		rsp.Errmsg = utils.RecodeText(utils.RECODE_NODATA)
		return err
	}
	if req.Text != rnd {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		//返回自定义的error数据
		return errors.New("验证码输入错误")
	}
	//如果成功,发送短信,存储短信验证码
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAI4FqQGHmaCcTAeZjoZ2im", "twt9iWtqIITPMUYJg0SnNmkGGxYpfg")
	if err != nil {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		return err
	}
	//获取6位数随机码
	myRnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06d", myRnd.Int31n(1000000))

	//初始化请求对象
	request := requests.NewCommonRequest()
	request.Method = "POST"                  //设置请求方法
	request.Scheme = "https"                 // https | http   //设置请求协议
	request.Domain = "dysmsapi.aliyuncs.com" //域名
	request.Version = "2017-05-25"           //版本号
	request.ApiName = "SendSms"              //api名称
	request.QueryParams["PhoneNumbers"] = req.Mobile
	request.QueryParams["SignName"] = "橘子"
	request.QueryParams["TemplateCode"] = "SMS_174272205"
	request.QueryParams["TemplateParam"] = `{"code":` + vcode + `}` //发送短信验证码

	response, err := client.ProcessCommonRequest(request)
	//如果不成功
	if !response.IsSuccess() {
		rsp.Errno = utils.RECODE_SMSERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_SMSERR)
		return errors.New("发送短信失败")
	}
	//存储短信验证码  存redis中
	//如果失败,直接就返回错误信息
	err = model.SaveSmsCode(req.Mobile, vcode)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}

	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	return nil

	return nil
}

//注册
//func(e *Register)MicroRegister(ctx context.Context,req *register.RegRequest,rsp *register.RegResponse) error{}
func (e *Register)MicroRegister(ctx context.Context, req *register.RegRequest, rsp *register.RegResponse) error {
	//实现具体的业务  把数据存储到mysql中  校验短信验证码是否正确
	//校验短信验证码会否正确
	smsCode, err := model.GetSmsCode(req.Mobile)
	if err != nil {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		return nil
	}
	if smsCode!=req.SmsCode{
		rsp.Errno=utils.RECODE_SMSERR
		rsp.Errmsg=utils.RecodeText(utils.RECODE_SMSERR)
		return  errors.New("验证码错误")
	}
	//存储用户数据到MySQL上
	//给密码加密
	pwdByte := sha256.Sum256([]byte(req.Password))
	pwd_hash := string(pwdByte[:])
	err = model.SaveUser(req.Mobile, pwd_hash)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	return nil

}
