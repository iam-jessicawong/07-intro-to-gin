package routers

import (
	"07-gin-get-started/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default() //kalau mau set router routes nya selalu pake gin.Default

	router.POST("/cars", controllers.CreateCar)
	router.PUT("/cars/:carID", controllers.UpdateCar)
	router.GET("/cars/:carID", controllers.GetCar)
	router.GET("/cars", controllers.GetCars)
	router.DELETE("/cars/:carID", controllers.DeleteCar)

	return router
}
