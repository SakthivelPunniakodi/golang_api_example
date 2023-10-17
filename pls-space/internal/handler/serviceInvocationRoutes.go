// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package handler

import (
	"sampleapi/pls-shared/constant"
	daprsvr "sampleapi/pls-shared/dapr/server"
)

func MapServiceInvocationRoutes(server daprsvr.Server, handler SpaceHandler) error {
	if err := server.AddServiceInvocationHandler(constant.AllocateSpace, handler.AllocateSpace); err != nil {
		return err
	}

	if err := server.AddServiceInvocationHandler(constant.DeallocateSpace, handler.DeallocateSpace); err != nil {
		return err
	}

	return nil
}
