package handler

import (
	"GinHello/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

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