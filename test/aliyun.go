package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

/*
func main() {
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
	request.Method = "POST"//设置请求方法
	request.Scheme = "https" // https | http   //设置请求协议
	request.Domain = "dysmsapi.aliyuncs.com"  //域名
	request.Version = "2017-05-25"			//版本号
	request.ApiName = "SendSms"				//api名称
	request.QueryParams["PhoneNumbers"] = req.Mobile
	request.QueryParams["SignName"] ="橘子"
	request.QueryParams["TemplateCode"]  = "SMS_174272205"
	request.QueryParams["TemplateParam"] = `{"code":`+vcode+`}`   //发送短信验证码

	response, err := client.ProcessCommonRequest(request)
	//如果不成功
	if !response.IsSuccess(){
		rsp.Errno = utils.RECODE_SMSERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_SMSERR)
		return errors.New("发送短信失败")
	}
	//存储短信验证码  存redis中
	err = model.SaveSmsCode(req.Mobile,vcode)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}

	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	return nil*/
