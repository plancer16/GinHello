package router

import (
	"GinHello/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../templates/*")
	} else {
		router.LoadHTMLGlob("templates/*")
	}
	router.Static("/statics","./statics")
	//router.GET("/", func(context *gin.Context) {
	//	context.String(http.StatusOK, "hello gin")
	//})
	//router.POST("/", func(context *gin.Context) {
	//	context.String(http.StatusOK, "hello gin post method")
	//})
	//router.GET("/", retHelloGinAndMethod)
	//router.GET("/user/:name", handler.UserSave)//http://localhost:8080/user/{name}
	//router.GET("/user",handler.UserSaveByQuery)//http://localhost:8080/user?name=lisi&age=20
	//router.POST("/", retHelloGinAndMethod)

	index := router.Group("/")
	{
		index.Any("", handler.Index)
	}

	userRouter := router.Group("/user")
	{
		userRouter.GET("/:name", handler.UserSave)  //从/{name}中获取
		userRouter.GET("", handler.UserSaveByQuery) //user?name=abc
	}

	return router
}

func retHelloGinAndMethod(context *gin.Context) {
	context.String(http.StatusOK, "hello gin "+strings.ToLower(context.Request.Method)+" method")
}
