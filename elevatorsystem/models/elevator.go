package models

import (
	"fmt"
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

func (e *Elevator) AddRequest(r *Request) bool {
	select {
	case e.requests <- <-e.requests:
		fmt.Printf("Lift : %d, From : %d ,TO:%d", e.id, r.sourceFloor, r.destinationFloor)
		return true

	default:
		return false
	}
}

func (e *Elevator) getCurrentFloor() int {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.currentFloor

}
