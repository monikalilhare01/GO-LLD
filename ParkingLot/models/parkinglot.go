package models

type ParkingLot struct {
	level []*Level
}

var instance *ParkingLot

func GetParkingLotInstance() *ParkingLot {
	if instance == nil {
		instance = &ParkingLot{level: []*Level{}}
	}
	return instance
}

func (p *ParkingLot) AddLevel(level *Level) {
	p.level = append(p.level, level)
}

func (p *ParkingLot) ParkVehicle(vehicle Vehicle) bool {
	for _, level := range p.level {
		if level.ParkVehicle(vehicle) {
			return true
		}
	}
	return false
}

func (p *ParkingLot) UnParkVehicle(vehicle Vehicle) bool {
	for _, level := range p.level {
		if level.UnParkVehicle(vehicle) {
			return true
		}
	}
	return false
}

func (p *ParkingLot) DisplayAvailability() {
	for _, level := range p.level {
		level.DisplayAvailability()
	}
}
