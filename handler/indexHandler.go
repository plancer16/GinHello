package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Index(context *gin.Context)  {
	context.HTML(http.StatusOK,"index.tmpl",gin.H{
		"title": "hello gin " + strings.ToLower(context.Request.Method) + " method",
	})//不返回string，返回html；指定到哪一个文件，title后面的便于该文件取出
}
