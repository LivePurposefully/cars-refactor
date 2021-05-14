package vehicleapi

import (
	"database/sql"
	"log"
	"strconv"
	"github.com/gin-contrib/cors"

	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"

	"github.com/LivePurposefully/cars-refactor/cars/pkg/models"
	"github.com/LivePurposefully/cars-refactor/cars/pkg/vehicledbconstants"
	"github.com/LivePurposefully/cars-refactor/cars/pkg/vehicledb"
)

var DB *sql.DB

func init() {
	DB = vehicledb.SetupPostgresDb()
	if DB == nil {
		log.Panic("Db is nil")
	}
}

func InitializeRouterWithConfiguration() *gin.Engine{
	router := gin.Default()
	// enable CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8082"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	return router
}

func InitializeVehicleRoutes(router *gin.Engine){
		//GET '/'  --> all cars
		router.GET("/cars", func(c *gin.Context){
			Index(c)
		})
	
		//POST '/cars'  --> create cars
		router.POST("/cars", func(c *gin.Context) {
			Create(c)
		})
	
		//GET '/cars/:carid'  --> get single car
		router.GET("/cars/:carid", func(c *gin.Context) {
			Show(c)
		})
	
		//PUT '/cars/:carid'  --> modify that single car
		router.PUT("/cars/:carid", func(c *gin.Context) {
			Update(c)
		})
	
		//DELETE '/cars/:carid'  --> delete that single car
		router.DELETE("/cars/:carid", func(c *gin.Context) {
			Destroy(c)
	
		})

		router.Run()
}

func Index(c *gin.Context) {
	results := make([]models.Vehicle, 0)

	var rows *sql.Rows
	var err error

	rows, err = DB.Query(vehicledbconstants.SelectAllCarsQuery)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	// count of records returned
	count := 0

	for rows.Next() {
		var obj models.Vehicle

		err = rows.Scan(&obj.Id, &obj.Make, &obj.Model, &obj.Year)

		// if the scan was successful, load the row
		if err == nil {
			results = append(results, obj)
			count++
		}
	}

	// show count of successfully processed rows
	log.Println("Rows returned: " + strconv.Itoa(count))

	if err = rows.Err(); err != nil {
		// Abnormal termination of the rows loop
		// close should be called automatically in this case
		log.Println(err)
	}

	c.JSON(200, results)
	//c.JSON(200, storeData)
}

func Create(c *gin.Context){
	var car models.Vehicle

	err := c.ShouldBindJSON(&car) //binds the input data into 'motor' var
	if err != nil {
		c.JSON(422, gin.H{"message": "Unprocessable entity!"})
		return
	}

	/*
				InsertCarQuery string = `INSERT INTO cars
				(make, model, year)
				VALUES($1, $2, $3)
			RETURNING id, make, model, year`
	*/

	err = DB.QueryRow(vehicledbconstants.InsertCarQuery, car.Make, car.Model, car.Year).Scan(
		&car.Id, &car.Make, &car.Model, &car.Year)

	if err != nil {
		c.JSON(500, gin.H{"message": "Internal server error!"})
	}

	/*
		car.Id = nextId
		nextId++

		storeData = append(storeData, car)
	*/
	c.JSON(200, car)
}

func Show (c *gin.Context){
	carid, err := strconv.Atoi(c.Param("carid"))
		if err != nil {
			c.JSON(404, gin.H{
				"message": "car id not valid",
			})
			return
		}

		var car models.Vehicle

		err = DB.QueryRow(vehicledbconstants.SelectCarQuery, carid).Scan(&car.Id,
			&car.Make, &car.Model, &car.Year)

		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(404, gin.H{"message": "Car not found!"})
				return
			}
			c.JSON(500, gin.H{"message": "Internal server error!"})
			return
		}

		/*

			var car Vehicle
			for i := 0; i < len(storeData); i++ {
				if storeData[i].Id == carid {
					car = storeData[i]
					break
				}
			}
			if car.Id == 0 {
				// the car was not found
				c.JSON(404, gin.H{
					"message": "car not found",
				})
				return
			}
		*/

		c.JSON(200, car)
}

func Update(c *gin.Context){
	carid, err := strconv.Atoi(c.Param("carid"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(404, gin.H{
			"message": "car id not valid",
		})
		return
	}

	var car models.Vehicle

	err = c.ShouldBindJSON(&car) //binds the input data into 'motor' var
	if err != nil {
		c.JSON(422, gin.H{"message": "Unprocessable entity!"})
		return
	}

	err = DB.QueryRow(vehicledbconstants.UpdateCarQuery, carid, car.Make, car.Model, car.Year).Scan(
		&car.Id, &car.Make, &car.Model, &car.Year)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"message": "Car not found!"})
			return
		}
		c.JSON(500, gin.H{"message": "Internal server error!"})
		return
	}

	/*
		found := false
		for i := 0; i < len(storeData); i++ {
			if storeData[i].Id == carid {
				car.Id = carid
				storeData[i] = car
				found = true
				break
			}
		}
		if !found {
			// the car was not found
			c.JSON(404, gin.H{
				"message": "car not found",
			})
			return
		}
	*/
	c.JSON(200, car)
}

func Destroy(c *gin.Context){
	carid, err := strconv.Atoi(c.Param("carid"))
	if err != nil {
		c.JSON(404, gin.H{
			"message": "car id not valid",
		})
		return
	}

	res, err := DB.Exec(vehicledbconstants.DeleteCarQuery, carid)
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal server error!"})
		return
	}
	count, err := res.RowsAffected()
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal server error!"})
		return
	}
	if count == 0 {
		c.JSON(404, gin.H{
			"message": "car not found",
		})
		return
	}

	/*
		found := false
		for i := 0; i < len(storeData); i++ {
			if storeData[i].Id == carid {
				storeData = append(storeData[:i], storeData[i+1:]...)
				found = true
				break
			}
		}

		if !found {
			c.JSON(404, gin.H{
				"message": "car not found",
			})
			return
		}
	*/
	c.JSON(200, gin.H{
		"message": "car successfully deleted",
	})
}