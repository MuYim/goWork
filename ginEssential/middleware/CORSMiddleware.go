package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//CORS跨域问题
func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//允许跨域的地址(想要访问本项目的地址【例如前端地址】)
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		//缓存时间
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		//可以跨域访问的方法 GET POST PUT DELETE....  *表示所有的方法都可以访问
		ctx.Writer.Header().Set("Access-Control_Allow-Methods", "*")
		//可以传输的头部信息
		ctx.Writer.Header().Set("Access-Control_Allow-Headers", "*")
		//是知传递证书
		ctx.Writer.Header().Set("Access-Control_Allow-Credentials", "true")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}
