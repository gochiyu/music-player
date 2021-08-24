package middleware

import (
	"Essential/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Failed(ctx, nil, fmt.Sprint(err))
			}
		}()
		ctx.Next()
	}
}
