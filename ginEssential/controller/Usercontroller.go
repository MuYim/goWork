package controller

import (
	"ginEssential/common"
	"ginEssential/dto"
	"ginEssential/model"
	"ginEssential/response"
	"ginEssential/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Register(context *gin.Context) {
	DB := common.GetDB()
	//获取参数(不能获取到通过json传的数据)
	//name := context.PostForm("name")
	//telephone := context.PostForm("telephone")
	//password := context.PostForm("password")

	//使用map获取请求参数
	//var requestMap=make(map[string]string)
	//json.NewDecoder(context.Request.Body).Decode(&requestMap)

	//使用结构体接收请求参数
	//var requestUser = model.User{}
	//json.NewDecoder(context.Request.Body).Decode(&requestUser)

	//gin的BInd获取参数(参数必须使用json传递)
	var requestUser = model.User{}
	context.Bind(&requestUser)

	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	//数据验证
	if len(telephone) != 11 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "密码必须大于6位")
		return
	}
	//如果名称没有传，生成一个10位的字符串
	if len(name) == 0 {
		name = utils.RandomString(10)
	}
	log.Println(name, telephone, password)
	//判断手机号是否存在
	if isTelephoneExitst(DB, telephone) {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "手机号已经注册")
		return
	}
	//创建用户
	//密码加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(context, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashPassword),
	}
	//坑1:必须传指针，因为要修改newUser的值，不传指针是值传递，无法修改最后使用的newUser，导致gorm.Model没有赋值，最终导致数据插入失败
	//db.Create(newUser)
	DB.Create(&newUser)

	//返回结果
	//response.Success(context, nil, "注册成功")

	//注册成功直接登录，不需用户在进行手动登录---发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(context, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error : %v", err)
		return
	}
	//返回结果
	response.Success(context, gin.H{"token": token}, "注册成功")
}

func Login(ctx *gin.Context) {
	DB := common.GetDB()
	//获取参数
	var requestUser = model.User{}
	ctx.Bind(&requestUser)
	telephone := requestUser.Telephone
	password := requestUser.Password
	//数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码必须大于6位")
		return
	}
	//判断手机号是否存在
	var user model.User
	DB.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "账号不存在")
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Fail(ctx, nil, "密码错误")
		return
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error : %v", err)
		return
	}
	//返回结果
	response.Success(ctx, gin.H{"token": token}, "登录成功")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	//TODO user.(model.User)类型断言
	response.Success(ctx, gin.H{"user": dto.ToUserDto(user.(model.User))}, "")
}

func isTelephoneExitst(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
