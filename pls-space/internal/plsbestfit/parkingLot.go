// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package plsbestfit

import (
	"container/heap"
	"fmt"
	"sync"

	"sampleapi/pls-shared/constant"
)

type ParkingLot struct {
	mu     *sync.Mutex
	Floors []*ParkingFloor
}

func NewParkingLot(floorCount, floorCapacity int, spotTypes ...[]constant.ParkingSpotType) *ParkingLot {
	if len(spotTypes) != floorCount {
		panic("Number of floor types must match the number of floors")
	}

	mu := sync.Mutex{}

	lot := &ParkingLot{
		Floors: make([]*ParkingFloor, floorCount),
		mu:     &mu,
	}

	for i := 0; i < floorCount; i++ {
		lot.Floors[i] = NewParkingFloor(&mu, i+1, floorCapacity, spotTypes[i]...)
	}

	return lot
}

func (o *ParkingLot) AllocateSpot(spotType constant.ParkingSpotType) (int, int, error) {
	for _, floor := range o.Floors {
		switch spotType {
		case constant.ParkingSpotTypeCompact:
			floorNumber, spotNumber, err := o.getParkingSpot(floor.FloorNumber, spotType, &floor.CompactAvailableSpots, &floor.CompactOccupiedSpots)
			if err != nil {
				continue
			}

			return floorNumber, spotNumber, nil
		case constant.ParkingSpotTypeLarge:
			floorNumber, spotNumber, err := o.getParkingSpot(floor.FloorNumber, spotType, &floor.LargeAvailableSpots, &floor.LargeOccupiedSpots)
			if err != nil {
				continue
			}

			return floorNumber, spotNumber, nil
		case constant.ParkingSpotTypeHandicapped:
			floorNumber, spotNumber, err := o.getParkingSpot(floor.FloorNumber, spotType, &floor.HandicappedAvailableSpots, &floor.HandicappedOccupiedSpots)
			if err != nil {
				continue
			}

			return floorNumber, spotNumber, nil
		case constant.ParkingSpotTypeEVChargeStation:
			floorNumber, spotNumber, err := o.getParkingSpot(floor.FloorNumber, spotType, &floor.EVChargeStationAvailableSpots, &floor.EVChargeStationOccupiedSpots)
			if err != nil {
				continue
			}

			return floorNumber, spotNumber, nil
		case constant.ParkingSpotTypeMotorcycle:
			floorNumber, spotNumber, err := o.getParkingSpot(floor.FloorNumber, spotType, &floor.MotorcycleAvailableSpots, &floor.MotorcycleOccupiedSpots)
			if err != nil {
				continue
			}

			return floorNumber, spotNumber, nil
		}
	}

	return 0, 0, fmt.Errorf("No available spot of type %s on parking lot", spotType)
}

func (o *ParkingLot) getParkingSpot(floorNumber int, spotType constant.ParkingSpotType, availableSpots *MinHeap, occupiedSpots *map[int]*ParkingSpot) (int, int, error) {
	if len(*availableSpots) == 0 && len(*occupiedSpots) == 0 {
		return 0, 0, fmt.Errorf("No available spot of type %s on floor %d", spotType, floorNumber)
	}

	if len(*availableSpots) == 0 {
		return 0, 0, fmt.Errorf("Parking floor %d is full", floorNumber)
	}

	for i := 0; i < len(*availableSpots); i++ {
		o.mu.Lock()

		parkingSpot := heap.Pop(availableSpots).(*ParkingSpot)
		parkingSpot.IsOccupied = true
		(*occupiedSpots)[parkingSpot.SpotNumber] = parkingSpot

		o.mu.Unlock()

		return floorNumber, parkingSpot.SpotNumber, nil
	}

	return 0, 0, fmt.Errorf("")
}

func (o *ParkingLot) DeallocateSpot(floorNumber, spotNumber int, spotType constant.ParkingSpotType) error {
	if floorNumber < 1 || floorNumber > len(o.Floors) {
		return fmt.Errorf("Invalid floor number: %d", floorNumber)
	}

	floor := o.Floors[floorNumber-1]

	switch spotType {
	case constant.ParkingSpotTypeCompact:
		o.deallocateSpot(spotNumber, &floor.CompactAvailableSpots, floor.CompactOccupiedSpots)
	case constant.ParkingSpotTypeLarge:
		o.deallocateSpot(spotNumber, &floor.LargeAvailableSpots, floor.LargeOccupiedSpots)
	case constant.ParkingSpotTypeHandicapped:
		o.deallocateSpot(spotNumber, &floor.HandicappedAvailableSpots, floor.HandicappedOccupiedSpots)
	case constant.ParkingSpotTypeEVChargeStation:
		o.deallocateSpot(spotNumber, &floor.EVChargeStationAvailableSpots, floor.EVChargeStationOccupiedSpots)
	case constant.ParkingSpotTypeMotorcycle:
		o.deallocateSpot(spotNumber, &floor.MotorcycleAvailableSpots, floor.MotorcycleOccupiedSpots)
	}

	return nil
}

func (o *ParkingLot) deallocateSpot(spotNumber int, availableSpots *MinHeap, occupiedSpots map[int]*ParkingSpot) {
	parkingSpot := NewParkingSpot(spotNumber, occupiedSpots[spotNumber].SpotType)
	parkingSpot.IsOccupied = false

	o.mu.Lock()
	defer o.mu.Unlock()

	heap.Push(availableSpots, parkingSpot)
	delete(occupiedSpots, spotNumber)
}
