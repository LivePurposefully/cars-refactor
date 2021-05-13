package main

import (
	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"

	"github.com/LivePurposefully/cars-refactor/cars/pkg/vehicleapi"
)

func main() {

	router := vehicleapi.InitializeRouterWithConfiguration()

	vehicleapi.InitializeVehicleRoutes(router)
}
