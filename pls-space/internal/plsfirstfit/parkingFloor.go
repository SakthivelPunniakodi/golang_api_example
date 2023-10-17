// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package plsfirstfit

import (
	"fmt"

	"sampleapi/pls-shared/constant"
)

type ParkingFloor struct {
	FloorNumber int
	Spots       []ParkingSpot
}

func NewParkingFloor(floorNumber, capacity int, spotTypes ...constant.ParkingSpotType) *ParkingFloor {
	if len(spotTypes) != capacity {
		panic("Number of slot types must match the floor capacity")
	}

	floor := &ParkingFloor{
		FloorNumber: floorNumber,
		Spots:       make([]ParkingSpot, capacity),
	}

	for i := 0; i < capacity; i++ {
		floor.Spots[i] = NewParkingSlot(i+1, spotTypes[i])
	}

	return floor
}

func (o *ParkingFloor) AllocateSlot(spotType constant.ParkingSpotType) (int, error) {
	for i, slot := range o.Spots {
		if !slot.IsOccupied && slot.SpotType == spotType {
			o.Spots[i].IsOccupied = true
			return slot.SpotNumber, nil
		}
	}
	return 0, fmt.Errorf("Parking floor %d is full for spot type %s", o.FloorNumber, spotType)
}

func (o *ParkingFloor) DeallocateSlot(spotNumber int) error {
	if spotNumber <= 0 || spotNumber > len(o.Spots) {
		return fmt.Errorf("Invalid spot number on Floor %d", o.FloorNumber)
	}

	if o.Spots[spotNumber-1].IsOccupied {
		o.Spots[spotNumber-1].IsOccupied = false
		return nil
	}

	return fmt.Errorf("Parking spot on Floor %d, Spot %d is not occupied", o.FloorNumber, spotNumber)
}
