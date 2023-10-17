// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package dto

import "sampleapi/pls-shared/constant"

type AllocateSpaceReq struct {
	SpotType constant.ParkingSpotType `json:"spotType"`
}
