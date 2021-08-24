package controller

import (
	"Essential/model"
	"Essential/repository"
	"Essential/response"
	"Essential/vo"
	"github.com/gin-gonic/gin"
	"strconv"
)

type IcategoryController interface {
	RestController
}
type CategoryController struct {
	Repository repository.CategoryRepository
}

func NewCategoryController() IcategoryController {
	repository := repository.NewCategoryRepository()
	repository.DB.AutoMigrate(model.Category{})
	return CategoryController{Repository: repository}
}
func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Failed(ctx, nil, "数据验证错误,分类名称必填")
		return
	}
	category, err := c.Repository.Create(requestCategory.Name)
	if err != nil {
		//response.Failed(ctx, nil, "创建失败")
		panic(err)
		return
	}
	response.Success(ctx, gin.H{"category": category}, "")
}
func (c CategoryController) Update(ctx *gin.Context) {
	//绑定body中的参数
	//var requestCategory model.Category
	//ctx.Bind(&requestCategory)
	//if requestCategory.Name==""{
	//	response.Failed(ctx,nil,"数据验证错误,分类名称必填")
	//}
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Failed(ctx, nil, "数据验证错误,分类名称必填")
		return
	}
	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	//var updateCategory model.Category
	//if c.DB.First(&updateCategory,categoryId).RecordNotFound(){
	//	response.Failed(ctx,nil,"分类不存在")
	//	return
	//}
	updateCategory, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Success(ctx, gin.H{"category": updateCategory}, "分类不存在")
	}
	//更新分类
	//map
	//struct
	//name value
	//c.DB.Model(&updateCategory).Update("name",requestCategory.Name)
	category, err := c.Repository.Update(*updateCategory, requestCategory.Name)
	if err != nil {
		panic(err)
	}
	response.Success(ctx, gin.H{"category": category}, "修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	category, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Success(ctx, gin.H{"category": category}, "分类不存在")
	}
	response.Success(ctx, gin.H{"category": category}, "")

}
func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取 path 参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	c.Repository.DeleteById(categoryId)

	response.Success(ctx, nil, "删除成功")
}
