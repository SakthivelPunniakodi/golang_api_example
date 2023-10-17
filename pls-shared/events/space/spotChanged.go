// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package connect

import (
	"encoding/json"
	"fmt"

	"sampleapi/pls-shared/constant"
	daprsvr "sampleapi/pls-shared/dapr/server"
	ierr "sampleapi/pls-shared/errors"
	"sampleapi/pls-shared/events/common"
)

const eventRouteSpotChanged common.EventRoute = "/spotChanged"

const eventTypeSpotChanged common.EventType = "SpotChanged"

type Floor struct {
	FloorNumber int                                 `json:"floorNumber"`
	SpotTypes   map[constant.ParkingSpotType][]Spot `json:"spots"`
}

type Spot struct {
	SpotNumber int `json:"spotNumber"`
}

type SpotChangedInput struct {
	Floors []Floor `json:"floors"`
}

type SpotChangedEvent struct {
	common.Event
	SpotChangedInput
}

func CreateSpotChangedEvent(input SpotChangedInput) SpotChangedEvent {
	return SpotChangedEvent{
		Event: common.Event{
			Type: eventTypeSpotChanged,
		},
		SpotChangedInput: input,
	}
}

func GetSpotChangedEvent(data []byte) (SpotChangedEvent, error) {
	event := SpotChangedEvent{}
	if err := json.Unmarshal(data, &event); err != nil {
		return event, ierr.WrapErrorf(err, ierr.Unknown, "json.Unmarshal")
	}

	return event, nil
}

func GetSpotChangedSubscription() daprsvr.Subscription {
	return daprsvr.Subscription{
		Topic: constant.Space,
		Route: string(eventRouteSpotChanged),
		Match: fmt.Sprintf(`event.data.type == "%s"`, eventTypeSpotChanged),
	}
}
