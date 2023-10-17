// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package ticket

import (
	"time"

	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/constant"

	"github.com/shopspring/decimal"
)

type GenerateTicketReq struct {
	VehicleNumber string                   `json:"vehicleNumber"`
	VehicleType   constant.VehicleType     `json:"vehicleType"`
	FloorNumber   int                      `json:"floorNumber"`
	SpotNumber    int                      `json:"spotNumber"`
	SpotType      constant.ParkingSpotType `json:"spotType"`
}

type GenerateTicketRes struct {
	TicketID      string                   `json:"ticketId"`
	VehicleNumber string                   `json:"vehicleNumber"`
	VehicleType   constant.VehicleType     `json:"vehicleType"`
	EntryTime     time.Time                `json:"entryTime"`
	ValidationKey string                   `json:"validationKey"`
	FloorNumber   int                      `json:"floorNumber"`
	SpotNumber    int                      `json:"spotNumber"`
	SpotType      constant.ParkingSpotType `json:"spotType"`
}

type GetTicketInfoReq struct {
	TicketID      string `json:"ticketId"`
	ValidationKey string `json:"validationKey"`
}

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

type DiscardTicketReq struct {
	TicketID      string `json:"ticketId"`
	ValidationKey string `json:"validationKey"`
}
