// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package plsbestfit

import (
	"container/heap"
	"sync"

	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/constant"
)

type ParkingFloor struct {
	mu                            *sync.Mutex
	FloorNumber                   int
	CompactAvailableSpots         MinHeap
	CompactOccupiedSpots          map[int]*ParkingSpot
	LargeAvailableSpots           MinHeap
	LargeOccupiedSpots            map[int]*ParkingSpot
	HandicappedAvailableSpots     MinHeap
	HandicappedOccupiedSpots      map[int]*ParkingSpot
	EVChargeStationAvailableSpots MinHeap
	EVChargeStationOccupiedSpots  map[int]*ParkingSpot
	MotorcycleAvailableSpots      MinHeap
	MotorcycleOccupiedSpots       map[int]*ParkingSpot
}

func NewParkingFloor(mu *sync.Mutex, floorNumber, capacity int, spotTypes ...constant.ParkingSpotType) *ParkingFloor {
	if len(spotTypes) != capacity {
		panic("Number of spot types must match the floor capacity")
	}

	compactAvailableSpots := make(MinHeap, 0, capacity)
	largeAvailableSpots := make(MinHeap, 0, capacity)
	handicappedAvailableSpots := make(MinHeap, 0, capacity)
	evChargeStationAvailableSpots := make(MinHeap, 0, capacity)
	motorcycleAvailableSpots := make(MinHeap, 0, capacity)

	for i := 1; i <= capacity; i++ {
		parkingSpot := NewParkingSpot(i, spotTypes[i-1])
		switch spotTypes[i-1] {
		case constant.ParkingSpotTypeCompact:
			compactAvailableSpots = append(compactAvailableSpots, parkingSpot)
		case constant.ParkingSpotTypeLarge:
			largeAvailableSpots = append(largeAvailableSpots, parkingSpot)
		case constant.ParkingSpotTypeHandicapped:
			handicappedAvailableSpots = append(handicappedAvailableSpots, parkingSpot)
		case constant.ParkingSpotTypeEVChargeStation:
			evChargeStationAvailableSpots = append(evChargeStationAvailableSpots, parkingSpot)
		case constant.ParkingSpotTypeMotorcycle:
			motorcycleAvailableSpots = append(motorcycleAvailableSpots, parkingSpot)
		}
	}

	heap.Init(&compactAvailableSpots)
	heap.Init(&largeAvailableSpots)
	heap.Init(&handicappedAvailableSpots)
	heap.Init(&evChargeStationAvailableSpots)
	heap.Init(&motorcycleAvailableSpots)

	return &ParkingFloor{
		mu:                            mu,
		FloorNumber:                   floorNumber,
		CompactAvailableSpots:         compactAvailableSpots,
		CompactOccupiedSpots:          make(map[int]*ParkingSpot),
		LargeAvailableSpots:           largeAvailableSpots,
		LargeOccupiedSpots:            make(map[int]*ParkingSpot),
		HandicappedAvailableSpots:     handicappedAvailableSpots,
		HandicappedOccupiedSpots:      make(map[int]*ParkingSpot),
		EVChargeStationAvailableSpots: evChargeStationAvailableSpots,
		EVChargeStationOccupiedSpots:  make(map[int]*ParkingSpot),
		MotorcycleAvailableSpots:      motorcycleAvailableSpots,
		MotorcycleOccupiedSpots:       make(map[int]*ParkingSpot),
	}
}
