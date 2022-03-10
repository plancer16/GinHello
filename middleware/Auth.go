package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		cookie, e := context.Request.Cookie("user_cookie")
		//每次接口请求需要权限时，都会刷新cookie时间
		if e == nil {
			context.SetCookie(cookie.Name,cookie.Value,1000,cookie.Path,cookie.Domain,cookie.Secure,cookie.HttpOnly)
			context.Next()
		} else {
			context.Abort()//终止请求
			context.HTML(http.StatusUnauthorized, "401.tmpl", nil)
		}
	}
}