package main

import "crm/router"

func main() {
	// 初始化路由
	r := router.Init()

	r.Run(":8080")
}
