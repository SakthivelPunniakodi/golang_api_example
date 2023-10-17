// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package plsfirstfit

import (
	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/constant"
)

type ParkingSpot struct {
	SpotNumber int
	SpotType   constant.ParkingSpotType
	IsOccupied bool
}

func NewParkingSlot(spotNumber int, spotType constant.ParkingSpotType) ParkingSpot {
	return ParkingSpot{
		SpotNumber: spotNumber,
		SpotType:   spotType,
		IsOccupied: false,
	}
}
