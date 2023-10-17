// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package dto

import "sampleapi/pls-shared/constant"

type AllocateSpotReq struct {
	VehicleNumber             string `json:"vehicleNumber"`
	VehicleType               string `json:"vehicleType"`
	IsNeedHandicappedSpot     bool   `json:"isNeedHandicappedSpot"`
	IsNeedEVChargeStationSpot bool   `json:"isNeedEVChargeStationSpot"`
	VehicleTypeEnum           constant.VehicleType
}
