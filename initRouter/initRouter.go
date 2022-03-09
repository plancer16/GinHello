package initRouter

import (
	"GinHello/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../templates/*")
	} else {
		router.LoadHTMLGlob("templates/*")
	}
	router.StaticFile("/favicon.ico","./favicon.ico")
	router.Static("/statics","./statics")


	index := router.Group("/")
	{
		index.Any("", handler.Index)
	}

	userRouter := router.Group("/user")
	{
		userRouter.GET("/:name", handler.UserSave)  //从/{name}中获取
		userRouter.GET("", handler.UserSaveByQuery) //user?name=abc
		userRouter.POST("/register", handler.UserRegister)
		userRouter.POST("/login",handler.UserLogin)
	}
	return router
}
