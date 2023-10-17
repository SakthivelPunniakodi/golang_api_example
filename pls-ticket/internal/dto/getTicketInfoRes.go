// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package dto

import (
	"time"

	"sampleapi/pls-shared/constant"

	"github.com/shopspring/decimal"
)

type GetTicketInfoRes struct {
	TicketID      string                   `json:"ticketId"`
	VehicleNumber string                   `json:"vehicleNumber"`
	VehicleType   constant.VehicleType     `json:"vehicleType"`
	EntryTime     time.Time                `json:"entryTime"`
	FloorNumber   int                      `json:"floorNumber"`
	SpotNumber    int                      `json:"spotNumber"`
	SpotType      constant.ParkingSpotType `json:"spotType"`
	Fees          decimal.Decimal          `json:"fees"`
}
