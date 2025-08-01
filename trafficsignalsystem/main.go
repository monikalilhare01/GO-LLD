package main

import (
	"time"

	"github.com/monikalilhare01/GO-LLD/trafficsignalsystem/models"
)

func main() {
	trafficController := models.GetTrafficController()

	road1 := models.NewRoad("road1", "MG Road")
	road2 := models.NewRoad("road2", "Okhla")
	road3 := models.NewRoad("road3", "Ring Road")
	road4 := models.NewRoad("road4", "By-Pass")

	// Create traffic lights
	trafficLight1 := models.NewTrafficLight("TL1", 6000, 3000, 9000)
	trafficLight2 := models.NewTrafficLight("TL2", 6000, 3000, 9000)
	trafficLight3 := models.NewTrafficLight("TL3", 6000, 3000, 9000)
	trafficLight4 := models.NewTrafficLight("TL4", 6000, 3000, 9000)

	road1.SetTrafficLight(*trafficLight1)
	road2.SetTrafficLight(*trafficLight2)
	road3.SetTrafficLight(*trafficLight3)
	road4.SetTrafficLight(*trafficLight4)

	trafficController.AddRoad(road1)
	trafficController.AddRoad(road2)
	trafficController.AddRoad(road3)
	trafficController.AddRoad(road4)

	go trafficController.StartTrafficControl()

	// Simulate an emergency after some time
	time.Sleep(10 * time.Second)
	trafficController.HandleEmergency("road1")
}
