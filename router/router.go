package router

import (
	"chal8/controllers"
	"chal8/middlewares"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("users")
	{
		userRouter.POST("/register", controllers.RegisterUser)
		userRouter.POST("/login", controllers.LoginUser)
	}

	productRouter := r.Group("products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.GET("/:productID", controllers.GetProduct)
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.PUT("/:productID", middlewares.ProductAuthorization(), controllers.UpdateProduct)
		productRouter.DELETE("/:productID", middlewares.ProductAuthorization(), controllers.DeleteProduct)
	}

	return r
}
