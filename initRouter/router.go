package initRouter

import (
	"back_end/handler"
	"back_end/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	//中间件
	//router := gin.Default()
	router:=gin.New()
	router.Use(middleware.Logger(),gin.Recovery())

	test := router.Group("/test")
	{
		test.GET("", handler.Test)
		test.POST("/insert",handler.Insert)
		test.GET("/findAll",handler.GetAll)
		test.DELETE("/delete/:id",handler.Delete)
		test.GET("/getOne/:id",handler.GetOne)
		test.POST("/update/:id",handler.Update)
	}
	

	api:=router.Group("/user")
	{
		api.POST("/login",handler.CreateJwt)
		api.POST("/register",handler.Register)
		api.GET("/findAll",middleware.Auth(),handler.GetAll)
	}

	bussinessapi:=router.Group("/bussiness")
	{
		bussinessapi.POST("/login",handler.BussinessCreateJwt)
		bussinessapi.POST("/register",handler.BussinessRegister)
		bussinessapi.POST("/updatePassW/:id",handler.BussinessUpdatePassW)
		bussinessapi.GET(("/getOneBussiness/:id"),handler.GetOneBussiness)
	}

	shopsApi:=router.Group("/shops")
	{
		shopsApi.GET("/search",handler.ShopSearch)
	}

	//swag集成
	//url := ginSwagger.URL("http://localhost:2333/swagger/doc.json")
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}

