// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package plsbestfit

import "sampleapi/pls-shared/constant"

type ParkingSpot struct {
	SpotNumber int
	SpotType   constant.ParkingSpotType
	IsOccupied bool
}

func NewParkingSpot(spotNumber int, spotType constant.ParkingSpotType) *ParkingSpot {
	return &ParkingSpot{
		SpotNumber: spotNumber,
		SpotType:   spotType,
		IsOccupied: false,
	}
}
