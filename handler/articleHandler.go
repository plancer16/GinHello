package handler

import (
	"GinHello/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)
// @Summary 提交新的文章内容
// @Id 1
// @Tags 文章
// @version 1.0
// @Accept application/x-json-stream
// @Param article body model.Article true "文章"
// @Success 200 object model.Result 成功后返回值
// @Failure 409 object model.Result 添加失败
// @Router /article [post]
func Insert(context *gin.Context) {
	article := model.Article{}
	var id = -1 //没插入文章的标志位
	if e := context.ShouldBindJSON(&article); e == nil {
		id = article.Insert()
	}
	context.JSON(http.StatusOK, gin.H{
		"id": id,
	})
	//gin.H表示json的map对
}
// @Summary 通过文章 id 获取单个文章内容
// @version 1.0
// @Accept application/x-json-stream
// @Param id path int true "id"
// @Success 200 object model.Result 成功后返回值
// @Router /article/{id} [get]
func GetOne(ctx *gin.Context) {
	id := ctx.Param("id")
	i, e := strconv.Atoi(id)
	if e != nil {
		log.Panicln("id 不是 int 类型, id 转换失败", e.Error())
	}
	article := model.Article{Id: i}
	art := article.FindById()
	ctx.JSON(http.StatusOK, gin.H{
		"article": art,
	})
}

func GetAll(ctx *gin.Context) {
	article := model.Article{}
	articles := article.FindAll()
	ctx.JSON(http.StatusOK,gin.H{
		"articles":articles,
	})
}

func DeleteOne(ctx *gin.Context) {
	id := ctx.Param("id")
	i, e := strconv.Atoi(id)
	if e != nil {
		log.Panicln("id 不是 int 类型, id 转换失败",e.Error())
	}
	article := model.Article{Id: i}
	article.DeleteOne()
}