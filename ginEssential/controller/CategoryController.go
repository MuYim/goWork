package controller

import (
	"ginEssential/model"
	"ginEssential/repository"
	"ginEssential/response"
	"ginEssential/vo"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	//DB *gorm.DB
	Repository repository.CateGoryRepository
}

func NewCategoryController() ICategoryController {
	//db := common.GetDB()
	//db.AutoMigrate(model.Category{})
	repository := repository.NewCategoryRespository()
	repository.DB.AutoMigrate(model.Category{})

	return CategoryController{Repository: repository}
}

func (c CategoryController) Creat(ctx *gin.Context) {
	//自己验证
	//var requestCategory model.Category
	//ctx.Bind(&requestCategory)
	//
	//if requestCategory.Name == "" {
	//	response.Fail(ctx, nil, "数据验证错误，分类名称必填")
	//	return
	//}

	//使用gin验证,通过binding注解可以一次验证多个参数
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}

	category, err := c.Repository.Create(requestCategory.Name)
	if err != nil {
		//response.Fail(ctx, nil, "创建失败，请重试")
		//通过RecoveryMiddleware中间件将错误信息返回给前端
		panic(err)
		return
	}

	response.Success(ctx, gin.H{"category": category}, "创建分类成功")
}

func (c CategoryController) Update(ctx *gin.Context) {
	//使用gin验证,通过binding注解可以一次验证多个参数
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}

	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	//c.DB.First(&updateCategory, categoryId).RecordNotFound()
	updateCategory, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}
	//更新分类
	//map
	//struct
	//name value
	//c.DB.Model(&updateCategory).Update("name", requestCategory.Name)
	category, err := c.Repository.Update(*updateCategory, requestCategory.Name)
	if err != nil {
		response.Fail(ctx, nil, "更新失败，请重试")
	}
	response.Success(ctx, gin.H{"category": category}, "修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	category, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	response.Success(ctx, gin.H{"category": category}, "查询成功")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	err := c.Repository.DeleteById(categoryId)
	if err != nil {
		response.Fail(ctx, nil, "删除失败，请重试")
		return
	}

	response.Success(ctx, nil, "删除成功")
}
