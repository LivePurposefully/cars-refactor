package main

import (
	"github.com/LivePurposefully/cars-refactor/cars/pkg/vehicleapi"
)

func main() {

	router := vehicleapi.InitializeRouterWithConfiguration()

	vehicleapi.InitializeVehicleRoutes(router)
}
