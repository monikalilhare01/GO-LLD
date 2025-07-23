package models

import "fmt"

type Level struct {
	floor       int
	parkingSpot []*ParkingSpot
}

func (l *Level) NewLevel(floor int, numSpot int) *Level {
	level := &Level{floor: floor}
	bikeSpot := int(float64(numSpot) * 0.50)
	carSpot := int(float64(numSpot) * 0.40)

	for i := 1; i <= bikeSpot; i++ {
		level.parkingSpot = append(level.parkingSpot, NewParkingSpot(i, MOTORCYCLE))
	}
	for i := bikeSpot + 1; i <= bikeSpot+carSpot; i++ {
		level.parkingSpot = append(level.parkingSpot, NewParkingSpot(i, CAR))
	}
	for i := bikeSpot + carSpot + 1; i <= numSpot; i++ {
		level.parkingSpot = append(level.parkingSpot, NewParkingSpot(i, TRUCK))
	}
	return level
}

func (l *Level) ParkVehicle(vehicle Vehicle) bool {
	for _, spot := range l.parkingSpot {
		if spot.IsAvailable() && spot.GetVehichleType() == vehicle.GetVehichleType() {
			spot.ParkVehicle(vehicle)
			return true
		}
	}
	return false
}

func (l *Level) UnParkVehicle(vehicle Vehicle) bool {
	for _, spot := range l.parkingSpot {
		if !spot.IsAvailable() && spot.GetVehichleType() == vehicle.GetVehichleType() {
			spot.UnParkVehicle(vehicle)
			return true
		}
	}
	return false
}

func (l *Level) DisplayAvailability() {
	for _, spot := range l.parkingSpot {
		status := "Available"
		if !spot.IsAvailable() {
			status = "Occupied"
		}
		fmt.Println("Level :", l.floor, "Spot :", spot.GetSpotNumber(), "Vehicle Type:", spot.GetVehichleType(), "status: ", status)
	}
}
