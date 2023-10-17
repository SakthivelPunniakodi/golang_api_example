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

func TestParkingFloorAllocateSlot(t *testing.T) {
	testCases := []struct {
		name         string
		floorNumber  int
		capacity     int
		spotTypes    []constant.ParkingSpotType
		allocateSpot constant.ParkingSpotType
		expectedSpot int
		expectedErr  error
	}{
		{
			name:        "Allocate spot successfully",
			floorNumber: 1,
			capacity:    3,
			spotTypes: []constant.ParkingSpotType{
				constant.ParkingSpotTypeCompact,
				constant.ParkingSpotTypeLarge,
				constant.ParkingSpotTypeMotorcycle,
			},
			allocateSpot: constant.ParkingSpotTypeLarge,
			expectedSpot: 2,
			expectedErr:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			floor := plsfirstfit.NewParkingFloor(testCase.floorNumber, testCase.capacity, testCase.spotTypes...)

			spotNumber, err := floor.AllocateSlot(testCase.allocateSpot)

			if spotNumber != testCase.expectedSpot {
				t.Errorf("For test case %s, expected spot number %d, but got %d", testCase.name, testCase.expectedSpot, spotNumber)
			}

			if (err != nil && testCase.expectedErr == nil) || (err == nil && testCase.expectedErr != nil) || (err != nil && err.Error() != testCase.expectedErr.Error()) {
				t.Errorf("For test case %s, expected error '%v', but got error '%v'", testCase.name, testCase.expectedErr, err)
			}
		})
	}
}

func TestParkingFloorDeallocateSlot(t *testing.T) {
	testCases := []struct {
		name           string
		floorNumber    int
		capacity       int
		spotTypes      []constant.ParkingSpotType
		allocateSpot   constant.ParkingSpotType
		deallocateSpot int
		expectedErr    error
	}{
		{
			name:        "Deallocate successfully",
			floorNumber: 1,
			capacity:    3,
			spotTypes: []constant.ParkingSpotType{
				constant.ParkingSpotTypeCompact,
				constant.ParkingSpotTypeLarge,
				constant.ParkingSpotTypeMotorcycle,
			},
			allocateSpot: constant.ParkingSpotTypeLarge,
			expectedErr:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			floor := plsfirstfit.NewParkingFloor(testCase.floorNumber, testCase.capacity, testCase.spotTypes...)

			spotNumber, _ := floor.AllocateSlot(testCase.allocateSpot)

			err := floor.DeallocateSlot(spotNumber)
			if (err != nil && testCase.expectedErr == nil) || (err == nil && testCase.expectedErr != nil) || (err != nil && err.Error() != testCase.expectedErr.Error()) {
				t.Errorf("For test case %s, expected error '%v', but got error '%v'", testCase.name, testCase.expectedErr, err)
			}
		})
	}
}
