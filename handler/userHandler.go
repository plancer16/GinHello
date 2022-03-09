package handler

import (
	"GinHello/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UserSave(context *gin.Context) {
	username := context.Param("name")
	context.String(http.StatusOK, "用户"+username+"已经保存")
}

func UserSaveByQuery(context *gin.Context) {
	username := context.Query("name")
	//age := context.Query("age")
	age := context.DefaultQuery("age", "24")
	context.String(http.StatusOK, "用户:"+username+",年龄:"+age+"已经保存")
}

func UserRegister(context *gin.Context)  {
	var user model.UserModel//模型绑定
	if err := context.ShouldBind(&user); err != nil {
		log.Println("err->", err.Error())//log日志
		context.String(http.StatusBadRequest, "输入的数据不合法")
	} else {
		log.Println("email", user.Email, "password", user.Password, "password again", user.PasswordAgain)
		context.Redirect(http.StatusMovedPermanently, "/")//修改为重定向状态码
	}
}