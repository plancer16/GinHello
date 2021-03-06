package handler

import (
	"GinHello/model"
	"GinHello/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
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
		context.SetCookie("user_cookie", string(u.Id),1000,"/","localhost",false,true)
		log.Println("登录成功", u.Email)
		context.HTML(http.StatusOK,"index.tmpl",gin.H{
			"email": u.Email,
			"id": 	 u.Id,
		})//把id给到html页面，便于之后context.query(id)查出来
		//context.String(http.StatusOK,"email",u.Email)
	}
}

func UserProfile(context *gin.Context)  {
	id := context.Query("id")
	var user model.UserModel
	i, err := strconv.Atoi(id)
	u, e := user.QueryById(i)
	if e != nil || err != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
	}
	context.HTML(http.StatusOK, "user_profile.tmpl", gin.H{
		"user": u,
	})
}

func UpdateUserProfile(context *gin.Context) {
	var user model.UserModel
	if err := context.ShouldBind(&user); err != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": err.Error(),
		})
		log.Panicln("绑定发生错误", err.Error())
	}
	file, e := context.FormFile("avatar-file")
	if e != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("文件上传错误", e.Error())
	}

	path := utils.RootPath()
	path = filepath.Join(path, "avatar")
	e = os.MkdirAll(path, os.ModePerm)
	if e != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("无法创建文件夹", e.Error())
	}
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	e = context.SaveUploadedFile(file, filepath.Join(path, fileName))
	if e != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("无法保存文件", e.Error())
	}
	avatarUrl := "/avatar/" + fileName
	user.Avatar = sql.NullString{String: avatarUrl}
	e = user.Update(user.Id)
	if e != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("数据无法更新", e.Error())
	}
	context.Redirect(http.StatusMovedPermanently, "/user/profile?id="+strconv.Itoa(user.Id))
}