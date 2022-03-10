package initRouter

import (
	"GinHello/handler"
	"GinHello/middleware"
	"GinHello/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	//router := gin.Default()
	router := gin.New()
	router.Use(middleware.Logger(), gin.Recovery()) //改了日志形式

	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../templates/*")
	} else {
		router.LoadHTMLGlob("templates/*")
	}
	router.StaticFile("/favicon.ico", "./favicon.ico")
	router.Static("/statics", "./statics")
	router.StaticFS("/avatar", http.Dir(utils.RootPath()+"avatar/"))

	index := router.Group("/")
	{
		index.Any("", handler.Index)
	}

	userRouter := router.Group("/user")
	{
		userRouter.GET("/:name", handler.UserSave)  //从/{name}中获取
		userRouter.GET("", handler.UserSaveByQuery) //user?name=abc
		userRouter.POST("/register", handler.UserRegister)
		userRouter.POST("/login", handler.UserLogin)
		userRouter.GET("/profile/", middleware.Auth(), handler.UserProfile)
		userRouter.POST("/update", middleware.Auth(), handler.UpdateUserProfile)
	}
	return router
}
