package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

type Stu struct {
	gorm.Model
	Name     string
	PassWord string
}
//多对多
type Teacher struct {

}

var GlobalDB *gorm.DB

func InitModel() {

	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/ihome?parseTime=true")
	if err != nil {
		fmt.Println("打开数据库失败", err)
		return
	}
	//连接池设置
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(30)
	db.DB().SetConnMaxLifetime(60 * 30)
	GlobalDB = db
	//设置表名为单数形式
	db.SingularTable(true)
	//自动迁移 在gorm中建表默认是负数形式
	db.AutoMigrate(new(Stu))
	fmt.Println("创建数据库stu成功")

}
func InitData() {
	var stu Stu
	stu.Name = "cat"
	stu.PassWord = "123456"
	if err := GlobalDB.Create(&stu).Error; err != nil {
	}
	fmt.Println(stu)
}
//查询数据
func SearchData() {
	var stu Stu
	stu.PassWord = "123"
	//查询所有取第一条

	if err := GlobalDB.Where("pass_word=?", "123456").First(&stu); err != nil {
		fmt.Println("查询错误", err)
		return
	}
	fmt.Println(stu)
}
//更新数据
func UpdataData(){
	var stu Stu
	stu.Name="dog"
	stu.PassWord="654321"
	//按条件跟新
/*if err:=	GlobalDB.Model(&stu).Where("name=?","dog").Update("pass_word","000000").Error;err!=nil{
		fmt.Println("更新密码失败",err)
	}*/
	GlobalDB.Model(&stu).Where("id=1")
	fmt.Println(stu)

}
//删除数据(unscoped确定将数据从数据库中删除)
func DeleteData (){
	//删除数据,软删除/逻辑删除  数据是无价的
	//(unscoped函数确定将数据从数据库中删除)
	var stu Stu
	if err:=GlobalDB.Where("id=1").Delete(&stu).Error;err!=nil{
		fmt.Println("删除失败",err)
		return
	}
}