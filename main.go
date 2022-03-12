package main

import (
	_ "GinHello/docs"
	"GinHello/initRouter"
	"github.com/gin-gonic/gin"
)

// @title Gin swagger
// @version 1.0
// @description Gin swagger 示例项目

// @contact.name 1922
// @contact.url https://plancer16.github.io
// @contact.email 1922@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// 主服务启动
func main() {
	gin.SetMode(gin.ReleaseMode)
	router := initRouter.SetupRouter()
	router.Run()
}