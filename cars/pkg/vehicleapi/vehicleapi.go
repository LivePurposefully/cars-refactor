package vehicleapi

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	results := make([]models.Vehicle, 0)

	var rows *sql.Rows
	var err error

	rows, err = db.Query(SelectAllCarsQuery)
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