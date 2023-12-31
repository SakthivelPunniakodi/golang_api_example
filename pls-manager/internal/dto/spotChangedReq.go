// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package dto

import (
	spaceEvents "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/events/space"
)

type SpotChangedReq struct {
	EventType string              `json:"eventType"`
	Floors    []spaceEvents.Floor `json:"floors"`
}
