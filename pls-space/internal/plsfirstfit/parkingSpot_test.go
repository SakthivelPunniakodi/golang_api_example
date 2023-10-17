// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package plsfirstfit_test

import (
	"testing"

	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/constant"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-space/internal/plsfirstfit"
)

func TestNewParkingSlot(t *testing.T) {
	testCases := []struct {
		spotNumber     int
		spotType       constant.ParkingSpotType
		expectedResult plsfirstfit.ParkingSpot
	}{
		{
			spotNumber: 1,
			spotType:   constant.ParkingSpotTypeCompact,
			expectedResult: plsfirstfit.ParkingSpot{
				SpotNumber: 1,
				SpotType:   constant.ParkingSpotTypeCompact,
				IsOccupied: false,
			},
		},
		{
			spotNumber: 2,
			spotType:   constant.ParkingSpotTypeLarge,
			expectedResult: plsfirstfit.ParkingSpot{
				SpotNumber: 2,
				SpotType:   constant.ParkingSpotTypeLarge,
				IsOccupied: false,
			},
		},
	}

	for _, testCase := range testCases {
		result := plsfirstfit.NewParkingSlot(testCase.spotNumber, testCase.spotType)

		if result != testCase.expectedResult {
			t.Errorf("For SpotNumber=%d, SpotType=%s, expected %+v, but got %+v",
				testCase.spotNumber, testCase.spotType, testCase.expectedResult, result)
		}
	}
}
