package initRouter

import (
	"GinHello/handler"
	"GinHello/middleware"
	"GinHello/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	//swagger文档路由
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

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

	articleRouter := router.Group("")
	{
		//网址url只能是get，post、delete通过postman，post要构造请求体json
		articleRouter.POST("/article",handler.Insert)
		articleRouter.GET("/article/:id",handler.GetOne)
		articleRouter.GET("articles",handler.GetAll)
		articleRouter.DELETE("/article/:id",handler.DeleteOne)
	}

	return router
}
