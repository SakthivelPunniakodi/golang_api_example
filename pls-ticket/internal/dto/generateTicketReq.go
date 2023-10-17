// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package dto

import "sampleapi/pls-shared/constant"

type GenerateTicketReq struct {
	VehicleNumber string                   `json:"vehicleNumber"`
	VehicleType   constant.VehicleType     `json:"vehicleType"`
	FloorNumber   int                      `json:"floorNumber"`
	SpotNumber    int                      `json:"spotNumber"`
	SpotType      constant.ParkingSpotType `json:"spotType"`
}
