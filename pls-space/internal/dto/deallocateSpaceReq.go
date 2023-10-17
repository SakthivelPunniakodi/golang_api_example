// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package dto

import "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/constant"

type DeallocateSpaceReq struct {
	FloorNumber int                      `json:"floorNumber"`
	SpotNumber  int                      `json:"spotNumber"`
	SpotType    constant.ParkingSpotType `json:"spotType"`
}
