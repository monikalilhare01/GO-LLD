package main

import (
	"github.com/monikalilhare01/GO-LLD/ParkingLot/models"
)

func main() {
	//config.InitDB()

	parkingLot := models.GetParkingLotInstance()
	// Create levels using the NewLevel method on a Level pointer
	parkingLot.AddLevel(new(models.Level).NewLevel(1, 100))
	parkingLot.AddLevel(new(models.Level).NewLevel(2, 80))

	car := &models.BaseVehicle{NumberPlate: "ABC123", VehicleType: models.CAR}
	truck := &models.BaseVehicle{NumberPlate: "XYZ789", VehicleType: models.TRUCK}
	motorcycle := &models.BaseVehicle{NumberPlate: "M1234", VehicleType: models.MOTORCYCLE}

	parkingLot.ParkVehicle(car)
	parkingLot.ParkVehicle(truck)
	parkingLot.ParkVehicle(motorcycle)

	//parkingLot.DisplayAvailability()

	//parkingLot.UnParkVehicle(motorcycle)

	parkingLot.DisplayAvailability()
}
