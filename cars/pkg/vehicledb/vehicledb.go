package vehicledb

import(
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/LivePurposefully/cars-refactor/cars/pkg/vehicledbconstants"
)

func SetupPostgresDb() *sql.DB{
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

		return db
}