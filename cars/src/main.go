package main

import (
	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"

	"github.com/LivePurposefully/cars-refactor/cars/pkg/vehicleapi"
)

func main() {

	router := vehicleapi.InitializeRouterWithConfiguration()

	//GET '/'  --> all cars
	router.GET("/cars", func(c *gin.Context){
		vehicleapi.Index(c)
	})

	//POST '/cars'  --> create cars
	router.POST("/cars", func(c *gin.Context) {
		vehicleapi.Create(c)
	})

	//GET '/cars/:carid'  --> get single car
	router.GET("/cars/:carid", func(c *gin.Context) {
		vehicleapi.Show(c)
	})

	//PUT '/cars/:carid'  --> modify that single car
	router.PUT("/cars/:carid", func(c *gin.Context) {
		vehicleapi.Update(c)
	})

	//DELETE '/cars/:carid'  --> delete that single car
	router.DELETE("/cars/:carid", func(c *gin.Context) {
		vehicleapi.Destroy(c)

	})

	router.Run()
}
