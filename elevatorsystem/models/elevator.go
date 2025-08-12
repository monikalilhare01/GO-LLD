package models

import (
	"sync"
)

type Elevator struct {
	id               int
	capacity         int
	currentFloor     int
	currentDirection int
	requests         chan *Request
	stopChan         chan struct{}
	mu               sync.RWMutex
}

func NewElevator(id, capacity int) *Elevator {
	return &Elevator{
		id:               id,
		capacity:         capacity,
		currentFloor:     1,
		currentDirection: int(DirectionUp),
		requests:         make(chan *Request, capacity),
		stopChan:         make(chan struct{}),
	}
}
