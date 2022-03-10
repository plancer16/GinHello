package main

import (
	"GinHello/initRouter"
)
// 主服务启动
func main() {
	router := initRouter.SetupRouter()
	router.Run()
}