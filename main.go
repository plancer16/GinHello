package main

import (
	"GinHello/router"
)
// 主服务启动
func main() {
	router := router.SetupRouter()
	router.Run()
}
