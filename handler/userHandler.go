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
	}
	passwordAgain := context.PostForm("password-again")
	//原本用字段校验，修改字段后，只能后端校验
	if passwordAgain != user.Password {
		context.String(http.StatusBadRequest, "两次密码不一致")//在网页前端显示错误提示
		log.Panicln("两次密码不一致")//程序中断
	}
	id := user.Save()
	log.Println("id is", id)
	context.Redirect(http.StatusMovedPermanently, "/")//修改为重定向状态码
}

func UserLogin(context *gin.Context)  {
	var user model.UserModel
	if e := context.Bind(&user); e != nil {
		log.Panicln("login 绑定错误", e.Error())
	}
	u := user.QueryByEmail()
	if u.Password == user.Password {
		log.Println("登录成功", u.Email)
		context.HTML(http.StatusOK,"index.tmpl",gin.H{
			"email": u.Email,
		})
		//context.String(http.StatusOK,"email",u.Email)
	}
}