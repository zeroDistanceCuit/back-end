package main

import (
	"back_end/initRouter"
)

// @title Zero-Distance
// @version 1.0
// @description Zero-Distance 项目后端

// @contact.name merlynr
// @contact.url https://blog.fanyan.vip
// @contact.email lcq1013962426@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:2333

func main() {
	// 可以在命令行启动服务的时候通过 -port=端口号 ，来指定 web 服务的端口号
	// 如果没有指定会使用默认的 8080
	//flag.StringVar(&port, "port", "8080", "The server listening port")
	//flag.Parse()

	router:=initRouter.SetupRouter()
	router.Run(":2333")
}