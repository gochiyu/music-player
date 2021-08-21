package middleware

import (
	"Essential/common"
	"Essential/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//用户认证保护路由
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		//validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "权限不足",
			})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "权限不足",
			})
			ctx.Abort()
			return

		}
		//验证通过后获取claim中的userid
		userID := claims.UserId
		DB := common.DB
		var user model.User
		DB.First(&user, userID)
		//用户
		if userID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "权限不足",
			})
			ctx.Abort()
			return
		}
		//用户存在,将user信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}

}
