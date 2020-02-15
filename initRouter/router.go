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
		test.POST("/insert",handler.Insert)
		test.GET("/findAll",handler.GetAll)
		test.DELETE("/delete/:id",handler.Delete)
		test.GET("/getOne/:id",handler.GetOne)
		test.POST("/update/:id",handler.Update)
	}
	return router
}

