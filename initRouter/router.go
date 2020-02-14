package initRouter

import (
	"back_end/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	//中间件
	router := gin.Default()

	test := router.Group("/api/test")
	{
		test.GET("/", handler.Test)
	}
	return router
}
