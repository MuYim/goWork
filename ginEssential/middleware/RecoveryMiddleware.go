package middleware

import (
	"fmt"
	"ginEssential/response"
	"github.com/gin-gonic/gin"
)

//可以使panic(err)的错误信息返回给前端
func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Fail(ctx, nil, fmt.Sprint(err))
			}
		}()
		ctx.Next()
	}
}
