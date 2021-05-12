package vehicledbconstants

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