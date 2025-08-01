package models

import (
	"sync"
	"time"
)

type TrafficController struct {
	roads map[string]*Road
	mu    sync.Mutex
}

var instance *TrafficController
var once sync.Once

func GetTrafficController() *TrafficController {
	once.Do(func() {
		instance = &TrafficController{roads: make(map[string]*Road)}
	})
	return instance
}
func (tc *TrafficController) AddRoad(road *Road) {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	tc.roads[road.ID] = road
}

func (tc *TrafficController) StartTrafficControl() {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	for _, road := range tc.roads {
		trafficLight := road.trafficLight
		go func(tl *TrafficLight) {
			for {
				time.Sleep(time.Duration(tl.RedDuration) * time.Millisecond)
				tl.ChangeSignal(Green)
				time.Sleep(time.Duration(tl.GreenDuration) * time.Millisecond)
				tl.ChangeSignal(Yellow)
				time.Sleep(time.Duration(tl.YellowDuration) * time.Millisecond)
				tl.ChangeSignal(Red)
			}
		}(&trafficLight)
	}
}

func (tc *TrafficController) HandleEmergency(roadID string) {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	if road, exist := tc.roads[roadID]; exist {
		trafficLight := road.trafficLight
		trafficLight.ChangeSignal(Green)
	}
}
