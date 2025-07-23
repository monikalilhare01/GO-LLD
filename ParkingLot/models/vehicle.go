package models

type VehicleType int

const (
	MOTORCYCLE VehicleType = iota
	CAR
	TRUCK
)

type Vehicle interface {
	GetNumberPlate() string
	GetVehichleType() VehicleType
}

type BaseVehicle struct {
	NumberPlate string
	VehicleType VehicleType
}

func (b *BaseVehicle) GetNumberPlate() string {
	return b.NumberPlate
}
func (b *BaseVehicle) GetVehichleType() VehicleType {
	return b.VehicleType
}
