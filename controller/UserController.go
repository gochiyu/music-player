package controller

import (
	"Essential/common"
	"Essential/dto"
	"Essential/model"
	"Essential/response"
	util "Essential/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Register(r *gin.Context) {
	DB := common.GetDB()
	//获取参数
	telephone := r.PostForm("telephone")
	password := r.PostForm("password")
	name := r.PostForm("name")
	//数据验证
	if len(telephone) != 11 {
		response.Response(r, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(r, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	//如果名称没有传 给一个随机的字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	//判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		response.Response(r, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		return
	}
	log.Println(name, telephone, password)
	//新建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(r, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)
	//返回结果
	response.Success(r, nil, "注册成功")
}

func Login(r *gin.Context) {
	DB := common.GetDB()
	//获取参数
	telephone := r.PostForm("telephone")
	password := r.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		response.Response(r, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(r, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	//判断手机号是否存在
	var user model.User
	DB.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(r, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		//r.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}
	//判断密码是否正确
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		response.Failed(r, nil, "密码错误")
		//r.JSON(http.StatusBadRequest,gin.H{"code": 400, "msg": "密码错误"})
		return
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(r, http.StatusInternalServerError, 500, nil, "系统异常")
		//r.JSON(http.StatusInternalServerError,gin.H{"code": 500, "msg": "系统异常"})
		log.Println("token generate error:%v", err)
		return
	}
	//r.JSON(200, gin.H{
	//	"code":200,
	//	"data":gin.H{"token":token},
	//	"message": "登陆成功",
	//})
	response.Success(r, gin.H{"token": token}, "登陆成功")

}
func Info(r *gin.Context) {
	user, _ := r.Get("user")
	response.Success(r, gin.H{"user": dto.ToUserDto(user.(model.User))}, "获取用户信息成功")
	//r.JSON(http.StatusOK,gin.H{"code":200, "data":gin.H{"user":dto.ToUserDto(user.(model.User))}})
}
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
