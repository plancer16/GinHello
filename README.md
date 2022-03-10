## Tutoria of Web Dev with Gin
[链接地址](https://youngxhui.top/categories/gin/)

### 集成swagger
1、下载
~~~
go get -u github.com/swaggo/swag/cmd/swag
~~~
2、添加路由，对ginSwagger和swaggerFiles进行sync下载
~~~
url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
~~~
3、接口，main.go添加注释

4、terminal中执行
~~~
swag init
~~~
启动项目，访问http://localhost:8080/swagger/index.html