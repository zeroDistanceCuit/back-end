package initRouter

import (
	"back_end/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	//seag集成
	url := ginSwagger.URL("http://localhost:2333/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}

