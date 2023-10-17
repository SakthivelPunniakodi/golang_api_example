// ---------------------------------------------------------------------------------------------------------------------
//  Copyrights (c) 2022 Dygisec PTE LTD - All Rights Reserved.
//  Unauthorized sharing or copying of Triton source code via any medium is strictly prohibited.
//  Proprietary and confidential.
//
//  Author: gayan@dygisec.com
//  Created On: 3/8/2022
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
	daprsvr "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/dapr/server"
	spaceEvents "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/events/space"
)

func MapTopicEventHandlerRoutes(server daprsvr.Server, handler Manager) error {
	if err := server.AddTopicEventHandler(spaceEvents.GetSpotChangedSubscription(), handler.SpotChanged); err != nil {
		return err
	}

	return nil
}
