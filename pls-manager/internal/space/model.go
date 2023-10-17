// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package space

import "sampleapi/pls-shared/constant"

type AllocateSpaceReq struct {
	SpotType constant.ParkingSpotType `json:"spotType"`
}

type AllocateSpaceRes struct {
	FloorNumber int                      `json:"floorNumber"`
	SpotNumber  int                      `json:"spotNumber"`
	SpotType    constant.ParkingSpotType `json:"spotType"`
}

type DeallocateSpaceReq struct {
	FloorNumber int                      `json:"floorNumber"`
	SpotNumber  int                      `json:"spotNumber"`
	SpotType    constant.ParkingSpotType `json:"spotType"`
}
