// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package plsfirstfit

import "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/constant"

type ParkingLot struct {
	Floors []*ParkingFloor
}

func NewParkingLot(floorCount, floorCapacity int, spotTypes ...[]constant.ParkingSpotType) *ParkingLot {
	if len(spotTypes) != floorCount {
		panic("Number of floor types must match the number of floors")
	}

	lot := &ParkingLot{
		Floors: make([]*ParkingFloor, floorCount),
	}

	for i := 0; i < floorCount; i++ {
		lot.Floors[i] = NewParkingFloor(i+1, floorCapacity, spotTypes[i]...)
	}

	return lot
}
