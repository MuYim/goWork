package controller

import (
	"ginEssential/common"
	"ginEssential/model"
	"ginEssential/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func Register(context *gin.Context) {
	DB := common.GetDB()
	//获取参数
	name := context.PostForm("name")
	telephone := context.PostForm("telephone")
	password := context.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码必须大于6位"})
		return
	}
	//如果名称没有传，生成一个10位的字符串
	if len(name) == 0 {
		name = utils.RandomString(10)
	}
	log.Println(name, telephone, password)
	//判断手机号是否存在
	if isTelephoneExitst(DB, telephone) {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号已经注册"})
		return
	}
	//创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	//坑1:必须传指针，因为要修改newUser的值，不传指针是值传递，无法修改最后使用的newUser，导致gorm.Model没有赋值，最终导致数据插入失败
	//db.Create(newUser)
	DB.Create(&newUser)
	//返回结果
	context.JSON(http.StatusOK, gin.H{"msg": "注册成功"})
}

func isTelephoneExitst(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
