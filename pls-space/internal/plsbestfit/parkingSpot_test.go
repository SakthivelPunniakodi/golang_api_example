package plsbestfit

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/constant"
)

func TestNewParkingSpot(t *testing.T) {
	testCases := []struct {
		spotNumber int
		spotType   constant.ParkingSpotType
		expected   *ParkingSpot
	}{
		{
			spotNumber: 1,
			spotType:   constant.ParkingSpotTypeCompact,
			expected: &ParkingSpot{
				SpotNumber: 1,
				SpotType:   constant.ParkingSpotTypeCompact,
				IsOccupied: false,
			},
		},
		{
			spotNumber: 2,
			spotType:   constant.ParkingSpotTypeLarge,
			expected: &ParkingSpot{
				SpotNumber: 2,
				SpotType:   constant.ParkingSpotTypeLarge,
				IsOccupied: false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("NewParkingSpot(%d, %s)", tc.spotNumber, tc.spotType), func(t *testing.T) {
			actual := NewParkingSpot(tc.spotNumber, tc.spotType)

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected: %+v, but got: %+v", tc.expected, actual)
			}
		})
	}
}
