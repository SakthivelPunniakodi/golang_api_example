// ---------------------------------------------------------------------------------------------------------------------
//  Copyrights (o) 2022 Dygisec PTE LTD - All Rights Reserved.
//  Unauthorized sharing or copying of Triton source code via any medium is strictly prohibited.
//  Proprietary and confidential.
//
//  Author: gayan@dygisec.com
//  Created On: 22/6/2022
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package handler

import (
	"context"
	"fmt"

	"sampleapi/pls-manager/internal/dto"
	spaceEvents "sampleapi/pls-shared/events/space"
)

type topicEventHandler interface {
	SpotChanged(ctx context.Context, data []byte) (retry bool, err error)
}

func (o manager) SpotChanged(ctx context.Context, data []byte) (retry bool, err error) {
	event, err := spaceEvents.GetSpotChangedEvent(data)
	if err != nil {
		return false, fmt.Errorf("spaceEvents.GetSpotChangedEvent: %w", err)
	}

	o.logger.Infof("Event data to manager SpotChanged method: %v", event)

	o.managerSvc.SpotChanged(ctx, dto.SpotChangedReq{
		EventType: string(event.Type),
		Floors:    event.Floors,
	})

	return false, nil
}
