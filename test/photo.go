package main

import (
	"github.com/afocus/captcha"
	"image/color"
	"image/png"
	"net/http"
)

func main() {
	//初始化实例对象
	cap:=captcha.New()
	//设置字符集
	if err:=cap.SetFont("./comic.ttf");err!=nil{
		panic(err.Error())
	}
	//设置验证码图片大小
	cap.SetSize(128.,64)
	//设置混淆程度
	cap.SetDisturbance(captcha.NORMAL)
	//设置字体颜色
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	//设置背景色  background
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
	//创建验证码图片
	http.HandleFunc("/r", func(w http.ResponseWriter, r *http.Request) {
		//第一个参数是验证码的位数,第二个参数是验证码的类型   cap自己生成随机数,返回给调用者
		img, str := cap.Create(6, captcha.NUM)
		png.Encode(w, img)
		println(str)
	})
	http.HandleFunc("/c", func(w http.ResponseWriter, r *http.Request) {
		str := r.URL.RawQuery
		//调用者生成随机数,传递给cap
		img := cap.CreateCustom(str)
		png.Encode(w, img)
	})

	http.ListenAndServe(":8002", nil)

}
