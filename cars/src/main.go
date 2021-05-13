package main

import (
	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	"github.com/LivePurposefully/cars-refactor/cars/pkg/vehicleapi"
)

func main() {

	router := gin.Default()
	// enable CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8082"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

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
