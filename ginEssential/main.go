package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name      string `gorm:type:varchar(20);not null`
	Telephone string `gorm:type:varchar(11);not null;unique`
	Password  string `gorm:size:255;not null`
}

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "root"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	//自动创建表
	db.AutoMigrate(&User{})
	return db
}
func main() {
	db := InitDB()
	defer db.Close()

	r := gin.Default()
	//用户注册
	r.POST("/api/auth/register", func(context *gin.Context) {
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
			name = RandomString(10)
		}
		log.Println(name, telephone, password)
		//判断手机号是否存在
		if isTelephoneExitst(db, telephone) {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号已经注册"})
			return
		}
		//创建用户
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		//坑1:必须传指针，因为要修改newUser的值，不传指针是值传递，无法修改最后使用的newUser，导致gorm.Model没有赋值，最终导致数据插入失败
		//db.Create(newUser)
		db.Create(&newUser)
		//返回结果
		context.JSON(http.StatusOK, gin.H{"msg": "注册成功"})
	})

	r.Run(":8081")
}

func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func isTelephoneExitst(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
