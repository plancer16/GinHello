package main

import (
	"GinHello/router"
)

func main() {
	router := router.SetupRouter()
	router.Run()
}
