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

### Docker部署
当前代码在13节有静态资源无法加载的问题，可通过官方repo的gin_docker分支来体验docker部署

下面总结下流程：

1、在服务器安装docker、配置docker端口号、配置镜像下载加速网址、启动docker

2、拉取mysql镜像，创建容器(制定宿主机和容器端口号映射)，进入mysql容器，创建database

3、本地项目build.bat编译生成linux的可执行文件，配置dockerfile(在服务器拉取镜像，并创建容器，用于将可执行文件复制到容器中执行)，编辑dockerfile(设置连接到服务器的地址:docker端口号，并映射服务器端口到golang项目启动端口)

总结：通过dockerfile，将生成的linux下可执行文件复制到服务器容器运行，连接到服务器mysql的3306，最后映射到mysql容器的3306