package models

type Road struct {
	ID           string
	name         string
	trafficLight TrafficLight
}

func NewRoad(id, name string) *Road {
	return &Road{
		ID:   id,
		name: name,
	}
}

func (r *Road) SetTrafficLight(tl TrafficLight) {
	r.trafficLight = tl
}
