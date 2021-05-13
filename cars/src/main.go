package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	"github.com/LivePurposefully/cars-refactor/cars/pkg/models"
	"github.com/LivePurposefully/cars-refactor/cars/pkg/vehicleapi"
)



const (
	SelectAllCarsQuery string = `SELECT id, make, model, year FROM cars ORDER BY id ASC`

	SelectCarQuery string = `SELECT id, make, model, year FROM cars WHERE id = $1`

	DeleteCarQuery string = `DELETE FROM cars WHERE id = $1`

	InsertCarQuery string = `INSERT INTO cars 
		(make, model, year) 
		VALUES($1, $2, $3)
    RETURNING id, make, model, year`

	UpdateCarQuery string = `UPDATE cars
		SET
			make  = COALESCE($2, make),
			model = COALESCE($3, model),
      		year  = COALESCE($4, year)
		WHERE
			id = $1 
		RETURNING
			id, make, model, year`
)

var length = 0 //to start at least 0 length and increase overtime
var storeData = make([]models.Vehicle, length)
var nextId = 1 // the next ID in the databas

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

	// set up database connection
	dbConnStr := fmt.Sprintf("user=postgres password=mysecretpassword dbname=postgres host=localhost port=49002 sslmode=disable")
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Panic(err)
	}

	// test the connection
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	//GET '/'  --> all cars
	router.GET("/cars", func(c *gin.Context){
		vehicleapi.Index(c, db)
	})

	//POST '/cars'  --> create cars
	router.POST("/cars", func(c *gin.Context) {
		vehicleapi.Create(c, db)
	})

	//GET '/cars/:carid'  --> get single car
	router.GET("/cars/:carid", func(c *gin.Context) {
		vehicleapi.Show(c, db)
	})

	//PUT '/cars/:carid'  --> modify that single car
	router.PUT("/cars/:carid", func(c *gin.Context) {
		vehicleapi.Update(c, db)
	})

	//DELETE '/cars/:carid'  --> delete that single car
	router.DELETE("/cars/:carid", func(c *gin.Context) {
		vehicleapi.Destroy(c, db)

	})

	router.Run()
}
