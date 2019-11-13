package handler

import (
	getImage "wfw-MVC/service/getImage/proto/getImage"
	"context"
	"github.com/afocus/captcha"
	"image/color"
	"wfw-MVC/service/getImage/model"
	"wfw-MVC/service/getImage/utils"
	"encoding/json"
	"fmt"
)

type GetImage struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetImage) MicroGetImg(ctx context.Context, req *getImage.Request, rsp *getImage.Response) error {
	//生成图片验证码
	cap := captcha.New()
	//设置字符集
	if err := cap.SetFont("comic.ttf"); err != nil {
		panic(err.Error())
	}
	//设置验证码图片大小
	cap.SetSize(128, 64)
	//设置混淆程度
	cap.SetDisturbance(captcha.NORMAL)
	//设置字体颜色
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	//设置背景色  background
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
	//生成图片验证码
	//rand.Seed(time.Now().UnixNano())
	img, rnd := cap.Create(4, captcha.NUM)
	//存储验证码
	err := model.SaveImageRnd(req.Uuid, rnd)

	fmt.Println("err================",err)
	fmt.Println("img",img)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}

	//传递图片信息给调用者
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	//json序列化
	imhJson,err:=json.Marshal(img)
	rsp.Data=imhJson
	return nil
}
