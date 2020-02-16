package main

import (
	"back_end/initRouter"
)

// @title Gin swagger
// @version 1.0
// @description Zero-Distance 项目后端

// @contact.name merlynr
// @contact.url https://blog.fanyan.vip
// @contact.email lcq1013962426@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:2333

func main() {

	router:=initRouter.SetupRouter()
	router.Run(":2333")
}