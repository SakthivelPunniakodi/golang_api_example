// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package dto

import (
	"time"

	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/constant"
)

type AllocateSpotRes struct {
	TicketID      string                   `json:"ticketId"`
	VehicleNumber string                   `json:"vehicleNumber"`
	VehicleType   string                   `json:"vehicleType"`
	EntryTime     time.Time                `json:"entryTime"`
	ValidationKey string                   `json:"validationKey"`
	FloorNumber   int                      `json:"floorNumber"`
	SpotNumber    int                      `json:"spotNumber"`
	SpotType      constant.ParkingSpotType `json:"spotType"`
}
