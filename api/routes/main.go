package routes

import (
	"github.com/gin-gonic/gin"
	controllers "numismatico/api/controllers"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	cedulaController := controllers.NewCedulaController()
	v1 := router.Group("cedulas")
	{
		v1.GET("", cedulaController.FindAll)
		v1.POST("", cedulaController.Create)
		v1.DELETE("/:id", cedulaController.Delete)
	}
	return v1
}
